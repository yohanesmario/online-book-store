// Code generated by mockery v2.43.2. DO NOT EDIT.

package base

import mock "github.com/stretchr/testify/mock"

// MockDBTransactionProvider is an autogenerated mock type for the DBTransactionProvider type
type MockDBTransactionProvider[TDBConnection interface{}] struct {
	mock.Mock
}

type MockDBTransactionProvider_Expecter[TDBConnection interface{}] struct {
	mock *mock.Mock
}

func (_m *MockDBTransactionProvider[TDBConnection]) EXPECT() *MockDBTransactionProvider_Expecter[TDBConnection] {
	return &MockDBTransactionProvider_Expecter[TDBConnection]{mock: &_m.Mock}
}

// Transaction provides a mock function with given fields: txBody
func (_m *MockDBTransactionProvider[TDBConnection]) Transaction(txBody func(IDBConnection[TDBConnection]) error) error {
	ret := _m.Called(txBody)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(IDBConnection[TDBConnection]) error) error); ok {
		r0 = rf(txBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBTransactionProvider_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type MockDBTransactionProvider_Transaction_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - txBody func(IDBConnection[TDBConnection]) error
func (_e *MockDBTransactionProvider_Expecter[TDBConnection]) Transaction(txBody interface{}) *MockDBTransactionProvider_Transaction_Call[TDBConnection] {
	return &MockDBTransactionProvider_Transaction_Call[TDBConnection]{Call: _e.mock.On("Transaction", txBody)}
}

func (_c *MockDBTransactionProvider_Transaction_Call[TDBConnection]) Run(run func(txBody func(IDBConnection[TDBConnection]) error)) *MockDBTransactionProvider_Transaction_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(IDBConnection[TDBConnection]) error))
	})
	return _c
}

func (_c *MockDBTransactionProvider_Transaction_Call[TDBConnection]) Return(_a0 error) *MockDBTransactionProvider_Transaction_Call[TDBConnection] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactionProvider_Transaction_Call[TDBConnection]) RunAndReturn(run func(func(IDBConnection[TDBConnection]) error) error) *MockDBTransactionProvider_Transaction_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// NewMockDBTransactionProvider creates a new instance of MockDBTransactionProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDBTransactionProvider[TDBConnection interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDBTransactionProvider[TDBConnection] {
	mock := &MockDBTransactionProvider[TDBConnection]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
