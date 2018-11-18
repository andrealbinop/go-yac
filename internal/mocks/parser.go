// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"

import mock "github.com/stretchr/testify/mock"

// Parser is an autogenerated mock type for the Parser type
type Parser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: _a0
func (_m *Parser) Parse(_a0 io.Reader) (map[string]interface{}, error) {
	ret := _m.Called(_a0)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(io.Reader) map[string]interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
