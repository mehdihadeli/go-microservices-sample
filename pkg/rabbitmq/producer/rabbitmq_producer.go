package producer

import (
	"context"
	"emperror.dev/errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/serializer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	messageHeader "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/message_header"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/producer"
	types2 "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/types"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/utils"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/producer/options"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/rabbitmq/types"
	"github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

type rabbitMQProducer struct {
	logger                  logger.Logger
	connection              types.IConnection
	eventSerializer         serializer.EventSerializer
	rabbitmqProducerOptions *options.RabbitMQProducerOptions
}

func NewRabbitMQProducer(connection types.IConnection, builderFunc func(builder *options.RabbitMQProducerOptionsBuilder), logger logger.Logger, eventSerializer serializer.EventSerializer) (producer.Producer, error) {
	builder := options.NewRabbitMQProducerOptionsBuilder()
	if builderFunc != nil {
		builderFunc(builder)
	}
	return &rabbitMQProducer{logger: logger, connection: connection, eventSerializer: eventSerializer, rabbitmqProducerOptions: builder.Build()}, nil
}

func (r *rabbitMQProducer) Publish(ctx context.Context, topicOrExchangeName string, message types2.IMessage, metadata core.Metadata) error {
	//https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go/publisher_confirms.go
	if r.connection == nil {
		return errors.New("connection is nil")
	}

	if r.connection.IsClosed() {
		return errors.New("connection is closed, wait for connection alive")
	}

	// create a unique channel on the connection and in the end close the channel
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	metadata = core.FromMetadata(metadata)

	if metadata.ExistsKey(messageHeader.MessageId) == false {
		metadata.SetValue(messageHeader.MessageId, message.GeMessageId())
	}

	if metadata.ExistsKey(messageHeader.CorrelationId) == false {
		cid := uuid.NewV4().String()
		metadata.SetValue(messageHeader.CorrelationId, cid)
		message.SetCorrelationId(cid)
	}

	metadata.SetValue(messageHeader.Name, utils.GetMessageName(message))

	serializedObj, err := r.eventSerializer.Serialize(message)
	if err != nil {
		return err
	}

	metadata.SetValue(messageHeader.Type, serializedObj.EventType)

	var exchange string

	if topicOrExchangeName != "" {
		exchange = topicOrExchangeName
	} else {
		exchange = utils.GetTopicOrExchangeName(message)
	}

	err = r.ensureExchange(channel, exchange)
	if err != nil {
		return err
	}

	if err := channel.Confirm(false); err != nil {
		return err
	}

	confirms := make(chan amqp091.Confirmation)
	channel.NotifyPublish(confirms)

	props := amqp091.Publishing{
		CorrelationId: message.GetCorrelationId(),
		MessageId:     message.GeMessageId(),
		Timestamp:     time.Now(),
		Headers:       core.MetadataToMap(metadata),
		Type:          serializedObj.EventType,
		ContentType:   serializedObj.ContentType,
		Body:          serializedObj.Data,
		DeliveryMode:  2,
	}

	err = channel.PublishWithContext(
		ctx,
		exchange,
		utils.GetRoutingKey(message),
		true,
		false,
		props,
	)
	if err != nil {
		return err
	}

	if confirmed := <-confirms; !confirmed.Ack {
		return errors.New("ack not confirmed")
	}

	return nil
}

func (r *rabbitMQProducer) ensureExchange(channel *amqp091.Channel, exchangeName string) error {
	err := channel.ExchangeDeclare(
		exchangeName,
		string(r.rabbitmqProducerOptions.ExchangeOptions.Type),
		r.rabbitmqProducerOptions.ExchangeOptions.Durable,
		r.rabbitmqProducerOptions.ExchangeOptions.AutoDelete,
		false,
		false,
		r.rabbitmqProducerOptions.ExchangeOptions.Args,
	)
	if err != nil {
		return err
	}

	return nil
}
