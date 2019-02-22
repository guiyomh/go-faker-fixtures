// Code generated by mockery v1.0.0. DO NOT EDIT.

package contract

import mock "github.com/stretchr/testify/mock"

// MockRegister is an autogenerated mock type for the Register type
type MockRegister struct {
	mock.Mock
}

// Denormalize provides a mock function with given fields: table, data
func (_m *MockRegister) Denormalize(table string, data map[string]map[string]interface{}) (Bager, error) {
	ret := _m.Called(table, data)

	var r0 Bager
	if rf, ok := ret.Get(0).(func(string, map[string]map[string]interface{}) Bager); ok {
		r0 = rf(table, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Bager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]map[string]interface{}) error); ok {
		r1 = rf(table, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
