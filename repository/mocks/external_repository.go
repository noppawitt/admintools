// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/noppawitt/admintools/model"

// ExternalRepository is an autogenerated mock type for the ExternalRepository type
type ExternalRepository struct {
	mock.Mock
}

// FindParameter provides a mock function with given fields: cs, name
func (_m *ExternalRepository) FindParameter(cs string, name string) ([]model.ExternalParameter, error) {
	ret := _m.Called(cs, name)

	var r0 []model.ExternalParameter
	if rf, ok := ret.Get(0).(func(string, string) []model.ExternalParameter); ok {
		r0 = rf(cs, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ExternalParameter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(cs, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStoredProcedure provides a mock function with given fields: cs
func (_m *ExternalRepository) FindStoredProcedure(cs string) ([]model.ExternalStoredProcedure, error) {
	ret := _m.Called(cs)

	var r0 []model.ExternalStoredProcedure
	if rf, ok := ret.Get(0).(func(string) []model.ExternalStoredProcedure); ok {
		r0 = rf(cs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ExternalStoredProcedure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindView provides a mock function with given fields: cs
func (_m *ExternalRepository) FindView(cs string) ([]model.ExternalView, error) {
	ret := _m.Called(cs)

	var r0 []model.ExternalView
	if rf, ok := ret.Get(0).(func(string) []model.ExternalView); ok {
		r0 = rf(cs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ExternalView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
