// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// URL provides a mock function with given fields: requestURI
func (_m *Backend) URL(requestURI string) string {
	ret := _m.Called(requestURI)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(requestURI)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}