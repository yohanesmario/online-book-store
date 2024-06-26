// Code generated by mockery v2.43.2. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	base "github.com/yohanesmario/online-book-store/internal/domain/repository/base"

	model "github.com/yohanesmario/online-book-store/internal/domain/model"
)

// MockIOrderRepository is an autogenerated mock type for the IOrderRepository type
type MockIOrderRepository[TDBConnection interface{}] struct {
	mock.Mock
}

type MockIOrderRepository_Expecter[TDBConnection interface{}] struct {
	mock *mock.Mock
}

func (_m *MockIOrderRepository[TDBConnection]) EXPECT() *MockIOrderRepository_Expecter[TDBConnection] {
	return &MockIOrderRepository_Expecter[TDBConnection]{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: order
func (_m *MockIOrderRepository[TDBConnection]) CreateOrder(order model.Order) (*model.Order, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Order) (*model.Order, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(model.Order) *model.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(model.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIOrderRepository_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type MockIOrderRepository_CreateOrder_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - order model.Order
func (_e *MockIOrderRepository_Expecter[TDBConnection]) CreateOrder(order interface{}) *MockIOrderRepository_CreateOrder_Call[TDBConnection] {
	return &MockIOrderRepository_CreateOrder_Call[TDBConnection]{Call: _e.mock.On("CreateOrder", order)}
}

func (_c *MockIOrderRepository_CreateOrder_Call[TDBConnection]) Run(run func(order model.Order)) *MockIOrderRepository_CreateOrder_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Order))
	})
	return _c
}

func (_c *MockIOrderRepository_CreateOrder_Call[TDBConnection]) Return(_a0 *model.Order, _a1 error) *MockIOrderRepository_CreateOrder_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIOrderRepository_CreateOrder_Call[TDBConnection]) RunAndReturn(run func(model.Order) (*model.Order, error)) *MockIOrderRepository_CreateOrder_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetOrdersByIDs provides a mock function with given fields: orderIDs
func (_m *MockIOrderRepository[TDBConnection]) GetOrdersByIDs(orderIDs []int32) ([]*model.Order, error) {
	ret := _m.Called(orderIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersByIDs")
	}

	var r0 []*model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func([]int32) ([]*model.Order, error)); ok {
		return rf(orderIDs)
	}
	if rf, ok := ret.Get(0).(func([]int32) []*model.Order); ok {
		r0 = rf(orderIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func([]int32) error); ok {
		r1 = rf(orderIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIOrderRepository_GetOrdersByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrdersByIDs'
type MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetOrdersByIDs is a helper method to define mock.On call
//   - orderIDs []int32
func (_e *MockIOrderRepository_Expecter[TDBConnection]) GetOrdersByIDs(orderIDs interface{}) *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection] {
	return &MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection]{Call: _e.mock.On("GetOrdersByIDs", orderIDs)}
}

func (_c *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection]) Run(run func(orderIDs []int32)) *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int32))
	})
	return _c
}

func (_c *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection]) Return(_a0 []*model.Order, _a1 error) *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection]) RunAndReturn(run func([]int32) ([]*model.Order, error)) *MockIOrderRepository_GetOrdersByIDs_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetOrdersByUserIDAndLastID provides a mock function with given fields: userID, lastID, limit
func (_m *MockIOrderRepository[TDBConnection]) GetOrdersByUserIDAndLastID(userID int32, lastID int32, limit int) ([]*model.Order, error) {
	ret := _m.Called(userID, lastID, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersByUserIDAndLastID")
	}

	var r0 []*model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int32, int32, int) ([]*model.Order, error)); ok {
		return rf(userID, lastID, limit)
	}
	if rf, ok := ret.Get(0).(func(int32, int32, int) []*model.Order); ok {
		r0 = rf(userID, lastID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int32, int32, int) error); ok {
		r1 = rf(userID, lastID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIOrderRepository_GetOrdersByUserIDAndLastID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrdersByUserIDAndLastID'
type MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetOrdersByUserIDAndLastID is a helper method to define mock.On call
//   - userID int32
//   - lastID int32
//   - limit int
func (_e *MockIOrderRepository_Expecter[TDBConnection]) GetOrdersByUserIDAndLastID(userID interface{}, lastID interface{}, limit interface{}) *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection] {
	return &MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection]{Call: _e.mock.On("GetOrdersByUserIDAndLastID", userID, lastID, limit)}
}

func (_c *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection]) Run(run func(userID int32, lastID int32, limit int)) *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32), args[1].(int32), args[2].(int))
	})
	return _c
}

func (_c *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection]) Return(_a0 []*model.Order, _a1 error) *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection]) RunAndReturn(run func(int32, int32, int) ([]*model.Order, error)) *MockIOrderRepository_GetOrdersByUserIDAndLastID_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// Use provides a mock function with given fields: txConnection
func (_m *MockIOrderRepository[TDBConnection]) Use(txConnection base.IDBConnection[TDBConnection]) IOrderRepository[TDBConnection] {
	ret := _m.Called(txConnection)

	if len(ret) == 0 {
		panic("no return value specified for Use")
	}

	var r0 IOrderRepository[TDBConnection]
	if rf, ok := ret.Get(0).(func(base.IDBConnection[TDBConnection]) IOrderRepository[TDBConnection]); ok {
		r0 = rf(txConnection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(IOrderRepository[TDBConnection])
		}
	}

	return r0
}

// MockIOrderRepository_Use_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Use'
type MockIOrderRepository_Use_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// Use is a helper method to define mock.On call
//   - txConnection base.IDBConnection[TDBConnection]
func (_e *MockIOrderRepository_Expecter[TDBConnection]) Use(txConnection interface{}) *MockIOrderRepository_Use_Call[TDBConnection] {
	return &MockIOrderRepository_Use_Call[TDBConnection]{Call: _e.mock.On("Use", txConnection)}
}

func (_c *MockIOrderRepository_Use_Call[TDBConnection]) Run(run func(txConnection base.IDBConnection[TDBConnection])) *MockIOrderRepository_Use_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(base.IDBConnection[TDBConnection]))
	})
	return _c
}

func (_c *MockIOrderRepository_Use_Call[TDBConnection]) Return(_a0 IOrderRepository[TDBConnection]) *MockIOrderRepository_Use_Call[TDBConnection] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOrderRepository_Use_Call[TDBConnection]) RunAndReturn(run func(base.IDBConnection[TDBConnection]) IOrderRepository[TDBConnection]) *MockIOrderRepository_Use_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// NewMockIOrderRepository creates a new instance of MockIOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIOrderRepository[TDBConnection interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIOrderRepository[TDBConnection] {
	mock := &MockIOrderRepository[TDBConnection]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
