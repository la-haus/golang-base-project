// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SampleRepository is an autogenerated mock type for the SampleRepository type
type SampleRepository struct {
	mock.Mock
}

// SampleRepositoryFunction provides a mock function with given fields: c
func (_m *SampleRepository) SampleRepositoryFunction(c context.Context) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
