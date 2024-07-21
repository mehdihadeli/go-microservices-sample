// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	appendResult "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/es/models/append_result"

	expectedStreamVersion "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/es/models/stream_version"

	metadata "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/core/metadata"

	mock "github.com/stretchr/testify/mock"

	models "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/es/models"

	readPosition "github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/es/models/stream_position/read_position"

	uuid "github.com/satori/go.uuid"
)

// AggregateStore is an autogenerated mock type for the AggregateStore type
type AggregateStore[T models.IHaveEventSourcedAggregate] struct {
	mock.Mock
}

type AggregateStore_Expecter[T models.IHaveEventSourcedAggregate] struct {
	mock *mock.Mock
}

func (_m *AggregateStore[T]) EXPECT() *AggregateStore_Expecter[T] {
	return &AggregateStore_Expecter[T]{mock: &_m.Mock}
}

// Exists provides a mock function with given fields: ctx, aggregateId
func (_m *AggregateStore[T]) Exists(ctx context.Context, aggregateId uuid.UUID) (bool, error) {
	ret := _m.Called(ctx, aggregateId)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (bool, error)); ok {
		return rf(ctx, aggregateId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(ctx, aggregateId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, aggregateId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateStore_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type AggregateStore_Exists_Call[T models.IHaveEventSourcedAggregate] struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregateId uuid.UUID
func (_e *AggregateStore_Expecter[T]) Exists(ctx interface{}, aggregateId interface{}) *AggregateStore_Exists_Call[T] {
	return &AggregateStore_Exists_Call[T]{Call: _e.mock.On("Exists", ctx, aggregateId)}
}

func (_c *AggregateStore_Exists_Call[T]) Run(run func(ctx context.Context, aggregateId uuid.UUID)) *AggregateStore_Exists_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *AggregateStore_Exists_Call[T]) Return(_a0 bool, _a1 error) *AggregateStore_Exists_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AggregateStore_Exists_Call[T]) RunAndReturn(run func(context.Context, uuid.UUID) (bool, error)) *AggregateStore_Exists_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: ctx, aggregateId
func (_m *AggregateStore[T]) Load(ctx context.Context, aggregateId uuid.UUID) (T, error) {
	ret := _m.Called(ctx, aggregateId)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (T, error)); ok {
		return rf(ctx, aggregateId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) T); ok {
		r0 = rf(ctx, aggregateId)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, aggregateId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateStore_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type AggregateStore_Load_Call[T models.IHaveEventSourcedAggregate] struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregateId uuid.UUID
func (_e *AggregateStore_Expecter[T]) Load(ctx interface{}, aggregateId interface{}) *AggregateStore_Load_Call[T] {
	return &AggregateStore_Load_Call[T]{Call: _e.mock.On("Load", ctx, aggregateId)}
}

func (_c *AggregateStore_Load_Call[T]) Run(run func(ctx context.Context, aggregateId uuid.UUID)) *AggregateStore_Load_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *AggregateStore_Load_Call[T]) Return(_a0 T, _a1 error) *AggregateStore_Load_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AggregateStore_Load_Call[T]) RunAndReturn(run func(context.Context, uuid.UUID) (T, error)) *AggregateStore_Load_Call[T] {
	_c.Call.Return(run)
	return _c
}

// LoadWithReadPosition provides a mock function with given fields: ctx, aggregateId, position
func (_m *AggregateStore[T]) LoadWithReadPosition(ctx context.Context, aggregateId uuid.UUID, position readPosition.StreamReadPosition) (T, error) {
	ret := _m.Called(ctx, aggregateId, position)

	if len(ret) == 0 {
		panic("no return value specified for LoadWithReadPosition")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, readPosition.StreamReadPosition) (T, error)); ok {
		return rf(ctx, aggregateId, position)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, readPosition.StreamReadPosition) T); ok {
		r0 = rf(ctx, aggregateId, position)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, readPosition.StreamReadPosition) error); ok {
		r1 = rf(ctx, aggregateId, position)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateStore_LoadWithReadPosition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithReadPosition'
type AggregateStore_LoadWithReadPosition_Call[T models.IHaveEventSourcedAggregate] struct {
	*mock.Call
}

// LoadWithReadPosition is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregateId uuid.UUID
//   - position readPosition.StreamReadPosition
func (_e *AggregateStore_Expecter[T]) LoadWithReadPosition(ctx interface{}, aggregateId interface{}, position interface{}) *AggregateStore_LoadWithReadPosition_Call[T] {
	return &AggregateStore_LoadWithReadPosition_Call[T]{Call: _e.mock.On("LoadWithReadPosition", ctx, aggregateId, position)}
}

func (_c *AggregateStore_LoadWithReadPosition_Call[T]) Run(run func(ctx context.Context, aggregateId uuid.UUID, position readPosition.StreamReadPosition)) *AggregateStore_LoadWithReadPosition_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(readPosition.StreamReadPosition))
	})
	return _c
}

func (_c *AggregateStore_LoadWithReadPosition_Call[T]) Return(_a0 T, _a1 error) *AggregateStore_LoadWithReadPosition_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AggregateStore_LoadWithReadPosition_Call[T]) RunAndReturn(run func(context.Context, uuid.UUID, readPosition.StreamReadPosition) (T, error)) *AggregateStore_LoadWithReadPosition_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function with given fields: aggregate, _a1, ctx
func (_m *AggregateStore[T]) Store(aggregate T, _a1 metadata.Metadata, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	ret := _m.Called(aggregate, _a1, ctx)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 *appendResult.AppendEventsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(T, metadata.Metadata, context.Context) (*appendResult.AppendEventsResult, error)); ok {
		return rf(aggregate, _a1, ctx)
	}
	if rf, ok := ret.Get(0).(func(T, metadata.Metadata, context.Context) *appendResult.AppendEventsResult); ok {
		r0 = rf(aggregate, _a1, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appendResult.AppendEventsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(T, metadata.Metadata, context.Context) error); ok {
		r1 = rf(aggregate, _a1, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateStore_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type AggregateStore_Store_Call[T models.IHaveEventSourcedAggregate] struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - aggregate T
//   - _a1 metadata.Metadata
//   - ctx context.Context
func (_e *AggregateStore_Expecter[T]) Store(aggregate interface{}, _a1 interface{}, ctx interface{}) *AggregateStore_Store_Call[T] {
	return &AggregateStore_Store_Call[T]{Call: _e.mock.On("Store", aggregate, _a1, ctx)}
}

func (_c *AggregateStore_Store_Call[T]) Run(run func(aggregate T, _a1 metadata.Metadata, ctx context.Context)) *AggregateStore_Store_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T), args[1].(metadata.Metadata), args[2].(context.Context))
	})
	return _c
}

func (_c *AggregateStore_Store_Call[T]) Return(_a0 *appendResult.AppendEventsResult, _a1 error) *AggregateStore_Store_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AggregateStore_Store_Call[T]) RunAndReturn(run func(T, metadata.Metadata, context.Context) (*appendResult.AppendEventsResult, error)) *AggregateStore_Store_Call[T] {
	_c.Call.Return(run)
	return _c
}

// StoreWithVersion provides a mock function with given fields: aggregate, _a1, expectedVersion, ctx
func (_m *AggregateStore[T]) StoreWithVersion(aggregate T, _a1 metadata.Metadata, expectedVersion expectedStreamVersion.ExpectedStreamVersion, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	ret := _m.Called(aggregate, _a1, expectedVersion, ctx)

	if len(ret) == 0 {
		panic("no return value specified for StoreWithVersion")
	}

	var r0 *appendResult.AppendEventsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(T, metadata.Metadata, expectedStreamVersion.ExpectedStreamVersion, context.Context) (*appendResult.AppendEventsResult, error)); ok {
		return rf(aggregate, _a1, expectedVersion, ctx)
	}
	if rf, ok := ret.Get(0).(func(T, metadata.Metadata, expectedStreamVersion.ExpectedStreamVersion, context.Context) *appendResult.AppendEventsResult); ok {
		r0 = rf(aggregate, _a1, expectedVersion, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appendResult.AppendEventsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(T, metadata.Metadata, expectedStreamVersion.ExpectedStreamVersion, context.Context) error); ok {
		r1 = rf(aggregate, _a1, expectedVersion, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AggregateStore_StoreWithVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreWithVersion'
type AggregateStore_StoreWithVersion_Call[T models.IHaveEventSourcedAggregate] struct {
	*mock.Call
}

// StoreWithVersion is a helper method to define mock.On call
//   - aggregate T
//   - _a1 metadata.Metadata
//   - expectedVersion expectedStreamVersion.ExpectedStreamVersion
//   - ctx context.Context
func (_e *AggregateStore_Expecter[T]) StoreWithVersion(aggregate interface{}, _a1 interface{}, expectedVersion interface{}, ctx interface{}) *AggregateStore_StoreWithVersion_Call[T] {
	return &AggregateStore_StoreWithVersion_Call[T]{Call: _e.mock.On("StoreWithVersion", aggregate, _a1, expectedVersion, ctx)}
}

func (_c *AggregateStore_StoreWithVersion_Call[T]) Run(run func(aggregate T, _a1 metadata.Metadata, expectedVersion expectedStreamVersion.ExpectedStreamVersion, ctx context.Context)) *AggregateStore_StoreWithVersion_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T), args[1].(metadata.Metadata), args[2].(expectedStreamVersion.ExpectedStreamVersion), args[3].(context.Context))
	})
	return _c
}

func (_c *AggregateStore_StoreWithVersion_Call[T]) Return(_a0 *appendResult.AppendEventsResult, _a1 error) *AggregateStore_StoreWithVersion_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AggregateStore_StoreWithVersion_Call[T]) RunAndReturn(run func(T, metadata.Metadata, expectedStreamVersion.ExpectedStreamVersion, context.Context) (*appendResult.AppendEventsResult, error)) *AggregateStore_StoreWithVersion_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewAggregateStore creates a new instance of AggregateStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAggregateStore[T models.IHaveEventSourcedAggregate](t interface {
	mock.TestingT
	Cleanup(func())
}) *AggregateStore[T] {
	mock := &AggregateStore[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
