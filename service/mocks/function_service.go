// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import repository "github.com/noppawitt/admintools/repository"

// FunctionService is an autogenerated mock type for the FunctionService type
type FunctionService struct {
	mock.Mock
}

// Repo provides a mock function with given fields:
func (_m *FunctionService) Repo() repository.FunctionRepository {
	ret := _m.Called()

	var r0 repository.FunctionRepository
	if rf, ok := ret.Get(0).(func() repository.FunctionRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.FunctionRepository)
		}
	}

	return r0
}