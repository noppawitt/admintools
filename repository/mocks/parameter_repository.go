// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/noppawitt/admintools/model"

// ParameterRepository is an autogenerated mock type for the ParameterRepository type
type ParameterRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: parameter
func (_m *ParameterRepository) Create(parameter *model.Parameter) error {
	ret := _m.Called(parameter)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Parameter) error); ok {
		r0 = rf(parameter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields:
func (_m *ParameterRepository) Find() ([]model.Parameter, error) {
	ret := _m.Called()

	var r0 []model.Parameter
	if rf, ok := ret.Get(0).(func() []model.Parameter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Parameter)
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

// FindByFunctionID provides a mock function with given fields: functionID
func (_m *ParameterRepository) FindByFunctionID(functionID int) ([]model.Parameter, error) {
	ret := _m.Called(functionID)

	var r0 []model.Parameter
	if rf, ok := ret.Get(0).(func(int) []model.Parameter); ok {
		r0 = rf(functionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Parameter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(functionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *ParameterRepository) FindOne(id int) (*model.Parameter, error) {
	ret := _m.Called(id)

	var r0 *model.Parameter
	if rf, ok := ret.Get(0).(func(int) *model.Parameter); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Parameter)
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

// Remove provides a mock function with given fields: id
func (_m *ParameterRepository) Remove(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: parameter
func (_m *ParameterRepository) Save(parameter *model.Parameter) error {
	ret := _m.Called(parameter)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Parameter) error); ok {
		r0 = rf(parameter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
