// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/noppawitt/admintools/model"

// ApplicationRepository is an autogenerated mock type for the ApplicationRepository type
type ApplicationRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: name, dbName
func (_m *ApplicationRepository) Count(name string, dbName string) (int, error) {
	ret := _m.Called(name, dbName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(name, dbName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, dbName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: application
func (_m *ApplicationRepository) Create(application *model.Application) error {
	ret := _m.Called(application)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Application) error); ok {
		r0 = rf(application)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields:
func (_m *ApplicationRepository) Find() ([]model.Application, error) {
	ret := _m.Called()

	var r0 []model.Application
	if rf, ok := ret.Get(0).(func() []model.Application); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Application)
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
func (_m *ApplicationRepository) FindOne(id int) (*model.Application, error) {
	ret := _m.Called(id)

	var r0 *model.Application
	if rf, ok := ret.Get(0).(func(int) *model.Application); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Application)
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

// Paginate provides a mock function with given fields: name, dbName, sortBy, direction, limit, offset
func (_m *ApplicationRepository) Paginate(name string, dbName string, sortBy string, direction string, limit int, offset int) ([]model.Application, error) {
	ret := _m.Called(name, dbName, sortBy, direction, limit, offset)

	var r0 []model.Application
	if rf, ok := ret.Get(0).(func(string, string, string, string, int, int) []model.Application); ok {
		r0 = rf(name, dbName, sortBy, direction, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, int, int) error); ok {
		r1 = rf(name, dbName, sortBy, direction, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *ApplicationRepository) Remove(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: application
func (_m *ApplicationRepository) Save(application *model.Application) error {
	ret := _m.Called(application)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Application) error); ok {
		r0 = rf(application)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}