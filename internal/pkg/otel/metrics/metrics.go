package metrics

// https://github.com/riferrei/otel-with-golang/blob/main/main.go
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go
// https://opentelemetry.io/docs/instrumentation/go/manual/#metrics

import (
	"context"
	"time"

	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/config/environment"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/http/customecho/contracts"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/logger"

	"emperror.dev/errors"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type OtelMetrics struct {
	config      *MetricsOptions
	logger      logger.Logger
	appMetrics  AppMetrics
	environment environment.Environment
	provider    *metric.MeterProvider
}

// NewOtelMetrics adds otel metrics
func NewOtelMetrics(
	config *MetricsOptions,
	logger logger.Logger,
	environment environment.Environment,
) (*OtelMetrics, error) {
	if config == nil {
		return nil, errors.New("metrics config can't be nil")
	}

	otelMetrics := &OtelMetrics{
		config:      config,
		logger:      logger,
		environment: environment,
	}

	resource, err := otelMetrics.newResource()
	if err != nil {
		return nil, errors.WrapIf(err, "failed to create resource")
	}

	appMetrics, err := otelMetrics.initMetrics(resource)
	if err != nil {
		return nil, err
	}

	otelMetrics.appMetrics = appMetrics

	return otelMetrics, nil
}

func (o *OtelMetrics) Shutdown(ctx context.Context) error {
	return o.provider.Shutdown(ctx)
}

func (o *OtelMetrics) newResource() (*resource.Resource, error) {
	// https://github.com/uptrace/uptrace-go/blob/master/example/otlp-traces/main.go#L49C1-L56C5
	resource, err := resource.New(
		context.Background(),
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOS(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceName(o.config.ServiceName),
			semconv.ServiceVersion(o.config.Version),
			attribute.String("environment", o.environment.GetEnvironmentName()),
			semconv.TelemetrySDKVersionKey.String("v1.21.0"), // semconv version
			semconv.TelemetrySDKLanguageGo,
		))

	return resource, err
}

func (o *OtelMetrics) initMetrics(
	resource *resource.Resource,
) (AppMetrics, error) {
	metricsExporter, err := o.configExporters()
	if err != nil {
		return nil, err
	}

	batchExporters := lo.Map(
		metricsExporter,
		func(item metric.Reader, index int) metric.Option {
			return metric.WithReader(item)
		},
	)

	// https://opentelemetry.io/docs/instrumentation/go/exporting_data/#resources
	// Resources are a special type of attribute that apply to all spans generated by a process
	opts := append(
		batchExporters,
		metric.WithResource(resource),
	)

	// otel library collects metrics and send this metrics to some exporter like console or prometheus
	provider := metric.NewMeterProvider(opts...)

	// Register our MeterProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetMeterProvider(provider)
	o.provider = provider

	appMeter := NewAppMeter(o.config.InstrumentationName)

	return appMeter, nil
}

func (o *OtelMetrics) configExporters() ([]metric.Reader, error) {
	ctx := context.Background()

	var exporters []metric.Reader

	// use some otel collector endpoints
	metricOpts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithTimeout(5 * time.Second),
		otlpmetricgrpc.WithInsecure(),
	}

	if !o.config.UseOTLP { //nolint:nestif

		if o.config.UseStdout {
			console, err := stdoutmetric.New()
			if err != nil {
				return nil, errors.WrapIf(
					err,
					"error creating console exporter",
				)
			}

			consoleMetricExporter := metric.NewPeriodicReader(
				console,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second))

			exporters = append(exporters, consoleMetricExporter)
		}

		if o.config.ElasticApmExporterOptions != nil {
			// https://www.elastic.co/guide/en/apm/guide/current/open-telemetry.html
			// https://www.elastic.co/guide/en/apm/guide/current/open-telemetry-direct.html#instrument-apps-otel
			// https://github.com/anilsenay/go-opentelemetry-examples/blob/elastic/cmd/main.go#L35
			metricOpts = append(
				metricOpts,
				otlpmetricgrpc.WithEndpoint(
					o.config.ElasticApmExporterOptions.OTLPEndpoint,
				),
				otlpmetricgrpc.WithHeaders(
					o.config.ElasticApmExporterOptions.OTLPHeaders,
				),
			)

			// send otel traces to jaeger builtin collector endpoint (default grpc port: 4317)
			// https://opentelemetry.io/docs/collector/
			exporter, err := otlpmetricgrpc.New(ctx, metricOpts...)
			if err != nil {
				return nil, errors.WrapIf(
					err,
					"failed to create otlpmetric exporter for elastic-apm",
				)
			}

			elasticApmExporter := metric.NewPeriodicReader(
				exporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second))

			exporters = append(exporters, elasticApmExporter)
		}

		if o.config.UptraceExporterOptions != nil {
			// https://github.com/uptrace/uptrace-go/blob/master/example/otlp-traces/main.go#L49C1-L56C5
			// https://uptrace.dev/get/opentelemetry-go.html#exporting-traces
			// https://uptrace.dev/get/opentelemetry-go.html#exporting-metrics
			metricOpts = append(
				metricOpts,
				otlpmetricgrpc.WithEndpoint(
					o.config.UptraceExporterOptions.OTLPEndpoint,
				),
				otlpmetricgrpc.WithHeaders(
					o.config.UptraceExporterOptions.OTLPHeaders,
				),
			)

			// send otel traces to jaeger builtin collector endpoint (default grpc port: 4317)
			// https://opentelemetry.io/docs/collector/
			exporter, err := otlpmetricgrpc.New(ctx, metricOpts...)
			if err != nil {
				return nil, errors.WrapIf(
					err,
					"failed to create otlpmetric exporter for uptrace",
				)
			}

			uptraceExporter := metric.NewPeriodicReader(
				exporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second))

			exporters = append(exporters, uptraceExporter)
		}
		if o.config.SignozExporterOptions != nil {
			// https://signoz.io/docs/instrumentation/golang/#instrumentation-of-a-sample-golang-application
			// https://signoz.io/blog/distributed-tracing-golang/
			metricOpts = append(
				metricOpts,
				otlpmetricgrpc.WithEndpoint(
					o.config.SignozExporterOptions.OTLPEndpoint,
				),
				otlpmetricgrpc.WithHeaders(
					o.config.SignozExporterOptions.OTLPHeaders,
				),
			)

			// send otel traces to jaeger builtin collector endpoint (default grpc port: 4317)
			// https://opentelemetry.io/docs/collector/
			exporter, err := otlpmetricgrpc.New(ctx, metricOpts...)
			if err != nil {
				return nil, errors.WrapIf(
					err,
					"failed to create otlpmetric exporter for signoz",
				)
			}

			signozExporter := metric.NewPeriodicReader(
				exporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second))

			exporters = append(exporters, signozExporter)
		} else {
			// https://prometheus.io/docs/prometheus/latest/getting_started/
			// https://prometheus.io/docs/guides/go-application/
			// prometheus exporter will collect otel metrics in prometheus registry
			// all prometheus exporters will add to a singleton `prometheus.DefaultRegisterer` registry in newConfig method and this registry will use via `promhttp.Handler` through http endpoint on `/metrics` and calls `Collect` on prometheus Reader interface inner signature prometheus.DefaultRegisterer
			prometheusExporter, err := prometheus.New()
			if err != nil {
				return nil, errors.WrapIf(
					err,
					"error creating prometheus exporter",
				)
			}
			exporters = append(exporters, prometheusExporter)
		}
	} else {
		for _, oltpProvider := range o.config.OTLPProviders {
			if !oltpProvider.Enabled {
				continue
			}

			metricOpts = append(metricOpts, otlpmetricgrpc.WithEndpoint(oltpProvider.OTLPEndpoint), otlpmetricgrpc.WithHeaders(oltpProvider.OTLPHeaders))

			// send otel metrics to an otel collector endpoint (default grpc port: 4317)
			// https://opentelemetry.io/docs/collector/
			// https://github.com/uptrace/uptrace-go/blob/master/example/otlp-metrics/main.go#L28
			// https://github.com/open-telemetry/opentelemetry-go/blob/main/exporters/otlp/otlpmetric/otlpmetricgrpc/example_test.go
			exporter, err := otlpmetricgrpc.New(ctx, metricOpts...)
			if err != nil {
				return nil, errors.WrapIf(err, "failed to create otlptracegrpc exporter")
			}
			metricExporter := metric.NewPeriodicReader(
				exporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second))

			exporters = append(exporters, metricExporter)
		}
	}

	return exporters, nil
}

// we could also use our existing server app port and a new /metrics endpoint instead of a new server with different port for our app metrics

func (o *OtelMetrics) RegisterMetricsEndpoint(
	server contracts.EchoHttpServer,
) {
	if o.config.UseOTLP {
		return
	}

	var metricsPath string
	if o.config.MetricsRoutePath == "" {
		metricsPath = "metrics"
	} else {
		metricsPath = o.config.MetricsRoutePath
	}

	// when we send request to /metrics endpoint, this handler gets singleton `prometheus.DefaultRegisterer` registry with calling `Collect` method on registered prometheus reader and get all get metrics and write them in /metrics endpoint output
	server.GetEchoInstance().
		GET(metricsPath, echo.WrapHandler(promhttp.Handler()))
}
