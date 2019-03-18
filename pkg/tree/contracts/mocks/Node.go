// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import contracts "github.com/guiyomh/charlatan/pkg/tree/contracts"
import mock "github.com/stretchr/testify/mock"

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// Add provides a mock function with given fields: node
func (_m *Node) Add(node contracts.Node) {
	_m.Called(node)
}

// EqualTo provides a mock function with given fields: other
func (_m *Node) EqualTo(other contracts.Node) bool {
	ret := _m.Called(other)

	var r0 bool
	if rf, ok := ret.Get(0).(func(contracts.Node) bool); ok {
		r0 = rf(other)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GreaterThan provides a mock function with given fields: other
func (_m *Node) GreaterThan(other contracts.Node) bool {
	ret := _m.Called(other)

	var r0 bool
	if rf, ok := ret.Get(0).(func(contracts.Node) bool); ok {
		r0 = rf(other)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Key provides a mock function with given fields:
func (_m *Node) Key() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LessThan provides a mock function with given fields: other
func (_m *Node) LessThan(other contracts.Node) bool {
	ret := _m.Called(other)

	var r0 bool
	if rf, ok := ret.Get(0).(func(contracts.Node) bool); ok {
		r0 = rf(other)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Maximum provides a mock function with given fields:
func (_m *Node) Maximum() contracts.Node {
	ret := _m.Called()

	var r0 contracts.Node
	if rf, ok := ret.Get(0).(func() contracts.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contracts.Node)
		}
	}

	return r0
}

// Minimum provides a mock function with given fields:
func (_m *Node) Minimum() contracts.Node {
	ret := _m.Called()

	var r0 contracts.Node
	if rf, ok := ret.Get(0).(func() contracts.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contracts.Node)
		}
	}

	return r0
}

// WalkBackward provides a mock function with given fields: iterator
func (_m *Node) WalkBackward(iterator contracts.Iterator) {
	_m.Called(iterator)
}

// WalkForward provides a mock function with given fields: iterator
func (_m *Node) WalkForward(iterator contracts.Iterator) {
	_m.Called(iterator)
}