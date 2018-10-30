// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/noppawitt/admintools/model"

// FunctionRepository is an autogenerated mock type for the FunctionRepository type
type FunctionRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: name, appName
func (_m *FunctionRepository) Count(name string, appName string) (int, error) {
	ret := _m.Called(name, appName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(name, appName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, appName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: function
func (_m *FunctionRepository) Create(function *model.Function) error {
	ret := _m.Called(function)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Function) error); ok {
		r0 = rf(function)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields:
func (_m *FunctionRepository) Find() ([]model.Function, error) {
	ret := _m.Called()

	var r0 []model.Function
	if rf, ok := ret.Get(0).(func() []model.Function); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *FunctionRepository) FindOne(id int) (*model.Function, error) {
	ret := _m.Called(id)

	var r0 *model.Function
	if rf, ok := ret.Get(0).(func(int) *model.Function); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneIncludeApplication provides a mock function with given fields: id
func (_m *FunctionRepository) FindOneIncludeApplication(id int) (*model.Function, error) {
	ret := _m.Called(id)

	var r0 *model.Function
	if rf, ok := ret.Get(0).(func(int) *model.Function); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: name, appName, sortBy, direction, limit, offset
func (_m *FunctionRepository) Paginate(name string, appName string, sortBy string, direction string, limit int, offset int) ([]model.Function, error) {
	ret := _m.Called(name, appName, sortBy, direction, limit, offset)

	var r0 []model.Function
	if rf, ok := ret.Get(0).(func(string, string, string, string, int, int) []model.Function); ok {
		r0 = rf(name, appName, sortBy, direction, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, int, int) error); ok {
		r1 = rf(name, appName, sortBy, direction, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *FunctionRepository) Remove(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: function
func (_m *FunctionRepository) Save(function *model.Function) error {
	ret := _m.Called(function)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Function) error); ok {
		r0 = rf(function)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
