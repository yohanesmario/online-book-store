// Code generated by mockery v2.43.2. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	base "github.com/yohanesmario/online-book-store/internal/domain/repository/base"

	model "github.com/yohanesmario/online-book-store/internal/domain/model"
)

// MockIUserRepository is an autogenerated mock type for the IUserRepository type
type MockIUserRepository[TDBConnection interface{}] struct {
	mock.Mock
}

type MockIUserRepository_Expecter[TDBConnection interface{}] struct {
	mock *mock.Mock
}

func (_m *MockIUserRepository[TDBConnection]) EXPECT() *MockIUserRepository_Expecter[TDBConnection] {
	return &MockIUserRepository_Expecter[TDBConnection]{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockIUserRepository[TDBConnection]) CreateUser(user model.User) (*model.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (*model.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockIUserRepository_CreateUser_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user model.User
func (_e *MockIUserRepository_Expecter[TDBConnection]) CreateUser(user interface{}) *MockIUserRepository_CreateUser_Call[TDBConnection] {
	return &MockIUserRepository_CreateUser_Call[TDBConnection]{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockIUserRepository_CreateUser_Call[TDBConnection]) Run(run func(user model.User)) *MockIUserRepository_CreateUser_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.User))
	})
	return _c
}

func (_c *MockIUserRepository_CreateUser_Call[TDBConnection]) Return(_a0 *model.User, _a1 error) *MockIUserRepository_CreateUser_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_CreateUser_Call[TDBConnection]) RunAndReturn(run func(model.User) (*model.User, error)) *MockIUserRepository_CreateUser_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockIUserRepository[TDBConnection]) GetUserByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockIUserRepository_GetUserByEmail_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockIUserRepository_Expecter[TDBConnection]) GetUserByEmail(email interface{}) *MockIUserRepository_GetUserByEmail_Call[TDBConnection] {
	return &MockIUserRepository_GetUserByEmail_Call[TDBConnection]{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *MockIUserRepository_GetUserByEmail_Call[TDBConnection]) Run(run func(email string)) *MockIUserRepository_GetUserByEmail_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIUserRepository_GetUserByEmail_Call[TDBConnection]) Return(_a0 *model.User, _a1 error) *MockIUserRepository_GetUserByEmail_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUserByEmail_Call[TDBConnection]) RunAndReturn(run func(string) (*model.User, error)) *MockIUserRepository_GetUserByEmail_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: id
func (_m *MockIUserRepository[TDBConnection]) GetUserByID(id int32) (*model.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int32) (*model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int32) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockIUserRepository_GetUserByID_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - id int32
func (_e *MockIUserRepository_Expecter[TDBConnection]) GetUserByID(id interface{}) *MockIUserRepository_GetUserByID_Call[TDBConnection] {
	return &MockIUserRepository_GetUserByID_Call[TDBConnection]{Call: _e.mock.On("GetUserByID", id)}
}

func (_c *MockIUserRepository_GetUserByID_Call[TDBConnection]) Run(run func(id int32)) *MockIUserRepository_GetUserByID_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32))
	})
	return _c
}

func (_c *MockIUserRepository_GetUserByID_Call[TDBConnection]) Return(_a0 *model.User, _a1 error) *MockIUserRepository_GetUserByID_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUserByID_Call[TDBConnection]) RunAndReturn(run func(int32) (*model.User, error)) *MockIUserRepository_GetUserByID_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// Use provides a mock function with given fields: txConnection
func (_m *MockIUserRepository[TDBConnection]) Use(txConnection base.IDBConnection[TDBConnection]) IUserRepository[TDBConnection] {
	ret := _m.Called(txConnection)

	if len(ret) == 0 {
		panic("no return value specified for Use")
	}

	var r0 IUserRepository[TDBConnection]
	if rf, ok := ret.Get(0).(func(base.IDBConnection[TDBConnection]) IUserRepository[TDBConnection]); ok {
		r0 = rf(txConnection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(IUserRepository[TDBConnection])
		}
	}

	return r0
}

// MockIUserRepository_Use_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Use'
type MockIUserRepository_Use_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// Use is a helper method to define mock.On call
//   - txConnection base.IDBConnection[TDBConnection]
func (_e *MockIUserRepository_Expecter[TDBConnection]) Use(txConnection interface{}) *MockIUserRepository_Use_Call[TDBConnection] {
	return &MockIUserRepository_Use_Call[TDBConnection]{Call: _e.mock.On("Use", txConnection)}
}

func (_c *MockIUserRepository_Use_Call[TDBConnection]) Run(run func(txConnection base.IDBConnection[TDBConnection])) *MockIUserRepository_Use_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(base.IDBConnection[TDBConnection]))
	})
	return _c
}

func (_c *MockIUserRepository_Use_Call[TDBConnection]) Return(_a0 IUserRepository[TDBConnection]) *MockIUserRepository_Use_Call[TDBConnection] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIUserRepository_Use_Call[TDBConnection]) RunAndReturn(run func(base.IDBConnection[TDBConnection]) IUserRepository[TDBConnection]) *MockIUserRepository_Use_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// NewMockIUserRepository creates a new instance of MockIUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIUserRepository[TDBConnection interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIUserRepository[TDBConnection] {
	mock := &MockIUserRepository[TDBConnection]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
