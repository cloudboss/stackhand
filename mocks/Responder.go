// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Responder is an autogenerated mock type for the Responder type
type Responder struct {
	mock.Mock
}

// FireFailed provides a mock function with given fields: reason
func (_m *Responder) FireFailed(reason string) error {
	ret := _m.Called(reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FireSuccess provides a mock function with given fields:
func (_m *Responder) FireSuccess() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
