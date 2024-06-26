// Code generated by mockery v2.43.2. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	base "github.com/yohanesmario/online-book-store/internal/domain/repository/base"

	model "github.com/yohanesmario/online-book-store/internal/domain/model"
)

// MockIBookRepository is an autogenerated mock type for the IBookRepository type
type MockIBookRepository[TDBConnection interface{}] struct {
	mock.Mock
}

type MockIBookRepository_Expecter[TDBConnection interface{}] struct {
	mock *mock.Mock
}

func (_m *MockIBookRepository[TDBConnection]) EXPECT() *MockIBookRepository_Expecter[TDBConnection] {
	return &MockIBookRepository_Expecter[TDBConnection]{mock: &_m.Mock}
}

// CreateBooks provides a mock function with given fields: books
func (_m *MockIBookRepository[TDBConnection]) CreateBooks(books []*model.Book) ([]*model.Book, error) {
	ret := _m.Called(books)

	if len(ret) == 0 {
		panic("no return value specified for CreateBooks")
	}

	var r0 []*model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.Book) ([]*model.Book, error)); ok {
		return rf(books)
	}
	if rf, ok := ret.Get(0).(func([]*model.Book) []*model.Book); ok {
		r0 = rf(books)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.Book) error); ok {
		r1 = rf(books)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIBookRepository_CreateBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBooks'
type MockIBookRepository_CreateBooks_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// CreateBooks is a helper method to define mock.On call
//   - books []*model.Book
func (_e *MockIBookRepository_Expecter[TDBConnection]) CreateBooks(books interface{}) *MockIBookRepository_CreateBooks_Call[TDBConnection] {
	return &MockIBookRepository_CreateBooks_Call[TDBConnection]{Call: _e.mock.On("CreateBooks", books)}
}

func (_c *MockIBookRepository_CreateBooks_Call[TDBConnection]) Run(run func(books []*model.Book)) *MockIBookRepository_CreateBooks_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*model.Book))
	})
	return _c
}

func (_c *MockIBookRepository_CreateBooks_Call[TDBConnection]) Return(_a0 []*model.Book, _a1 error) *MockIBookRepository_CreateBooks_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIBookRepository_CreateBooks_Call[TDBConnection]) RunAndReturn(run func([]*model.Book) ([]*model.Book, error)) *MockIBookRepository_CreateBooks_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetBooksByIDs provides a mock function with given fields: bookIds
func (_m *MockIBookRepository[TDBConnection]) GetBooksByIDs(bookIds []int32) ([]*model.Book, error) {
	ret := _m.Called(bookIds)

	if len(ret) == 0 {
		panic("no return value specified for GetBooksByIDs")
	}

	var r0 []*model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func([]int32) ([]*model.Book, error)); ok {
		return rf(bookIds)
	}
	if rf, ok := ret.Get(0).(func([]int32) []*model.Book); ok {
		r0 = rf(bookIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func([]int32) error); ok {
		r1 = rf(bookIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIBookRepository_GetBooksByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooksByIDs'
type MockIBookRepository_GetBooksByIDs_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetBooksByIDs is a helper method to define mock.On call
//   - bookIds []int32
func (_e *MockIBookRepository_Expecter[TDBConnection]) GetBooksByIDs(bookIds interface{}) *MockIBookRepository_GetBooksByIDs_Call[TDBConnection] {
	return &MockIBookRepository_GetBooksByIDs_Call[TDBConnection]{Call: _e.mock.On("GetBooksByIDs", bookIds)}
}

func (_c *MockIBookRepository_GetBooksByIDs_Call[TDBConnection]) Run(run func(bookIds []int32)) *MockIBookRepository_GetBooksByIDs_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int32))
	})
	return _c
}

func (_c *MockIBookRepository_GetBooksByIDs_Call[TDBConnection]) Return(_a0 []*model.Book, _a1 error) *MockIBookRepository_GetBooksByIDs_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIBookRepository_GetBooksByIDs_Call[TDBConnection]) RunAndReturn(run func([]int32) ([]*model.Book, error)) *MockIBookRepository_GetBooksByIDs_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// GetBooksByLastID provides a mock function with given fields: lastID, limit
func (_m *MockIBookRepository[TDBConnection]) GetBooksByLastID(lastID int32, limit int) ([]*model.Book, error) {
	ret := _m.Called(lastID, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetBooksByLastID")
	}

	var r0 []*model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(int32, int) ([]*model.Book, error)); ok {
		return rf(lastID, limit)
	}
	if rf, ok := ret.Get(0).(func(int32, int) []*model.Book); ok {
		r0 = rf(lastID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(int32, int) error); ok {
		r1 = rf(lastID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIBookRepository_GetBooksByLastID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooksByLastID'
type MockIBookRepository_GetBooksByLastID_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// GetBooksByLastID is a helper method to define mock.On call
//   - lastID int32
//   - limit int
func (_e *MockIBookRepository_Expecter[TDBConnection]) GetBooksByLastID(lastID interface{}, limit interface{}) *MockIBookRepository_GetBooksByLastID_Call[TDBConnection] {
	return &MockIBookRepository_GetBooksByLastID_Call[TDBConnection]{Call: _e.mock.On("GetBooksByLastID", lastID, limit)}
}

func (_c *MockIBookRepository_GetBooksByLastID_Call[TDBConnection]) Run(run func(lastID int32, limit int)) *MockIBookRepository_GetBooksByLastID_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32), args[1].(int))
	})
	return _c
}

func (_c *MockIBookRepository_GetBooksByLastID_Call[TDBConnection]) Return(_a0 []*model.Book, _a1 error) *MockIBookRepository_GetBooksByLastID_Call[TDBConnection] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIBookRepository_GetBooksByLastID_Call[TDBConnection]) RunAndReturn(run func(int32, int) ([]*model.Book, error)) *MockIBookRepository_GetBooksByLastID_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// Use provides a mock function with given fields: txConnection
func (_m *MockIBookRepository[TDBConnection]) Use(txConnection base.IDBConnection[TDBConnection]) IBookRepository[TDBConnection] {
	ret := _m.Called(txConnection)

	if len(ret) == 0 {
		panic("no return value specified for Use")
	}

	var r0 IBookRepository[TDBConnection]
	if rf, ok := ret.Get(0).(func(base.IDBConnection[TDBConnection]) IBookRepository[TDBConnection]); ok {
		r0 = rf(txConnection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(IBookRepository[TDBConnection])
		}
	}

	return r0
}

// MockIBookRepository_Use_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Use'
type MockIBookRepository_Use_Call[TDBConnection interface{}] struct {
	*mock.Call
}

// Use is a helper method to define mock.On call
//   - txConnection base.IDBConnection[TDBConnection]
func (_e *MockIBookRepository_Expecter[TDBConnection]) Use(txConnection interface{}) *MockIBookRepository_Use_Call[TDBConnection] {
	return &MockIBookRepository_Use_Call[TDBConnection]{Call: _e.mock.On("Use", txConnection)}
}

func (_c *MockIBookRepository_Use_Call[TDBConnection]) Run(run func(txConnection base.IDBConnection[TDBConnection])) *MockIBookRepository_Use_Call[TDBConnection] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(base.IDBConnection[TDBConnection]))
	})
	return _c
}

func (_c *MockIBookRepository_Use_Call[TDBConnection]) Return(_a0 IBookRepository[TDBConnection]) *MockIBookRepository_Use_Call[TDBConnection] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIBookRepository_Use_Call[TDBConnection]) RunAndReturn(run func(base.IDBConnection[TDBConnection]) IBookRepository[TDBConnection]) *MockIBookRepository_Use_Call[TDBConnection] {
	_c.Call.Return(run)
	return _c
}

// NewMockIBookRepository creates a new instance of MockIBookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIBookRepository[TDBConnection interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIBookRepository[TDBConnection] {
	mock := &MockIBookRepository[TDBConnection]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
