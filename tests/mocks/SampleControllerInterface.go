// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SampleControllerInterface is an autogenerated mock type for the SampleControllerInterface type
type SampleControllerInterface struct {
	mock.Mock
}

// GetSample provides a mock function with given fields: c
func (_m *SampleControllerInterface) GetSample(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}