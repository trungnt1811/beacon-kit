// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	mock "github.com/stretchr/testify/mock"
)

// Deposit is an autogenerated mock type for the Deposit type
type Deposit struct {
	mock.Mock
}

type Deposit_Expecter struct {
	mock *mock.Mock
}

func (_m *Deposit) EXPECT() *Deposit_Expecter {
	return &Deposit_Expecter{mock: &_m.Mock}
}

// GetIndex provides a mock function with given fields:
func (_m *Deposit) GetIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Deposit_GetIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIndex'
type Deposit_GetIndex_Call struct {
	*mock.Call
}

// GetIndex is a helper method to define mock.On call
func (_e *Deposit_Expecter) GetIndex() *Deposit_GetIndex_Call {
	return &Deposit_GetIndex_Call{Call: _e.mock.On("GetIndex")}
}

func (_c *Deposit_GetIndex_Call) Run(run func()) *Deposit_GetIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Deposit_GetIndex_Call) Return(_a0 math.U64) *Deposit_GetIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_GetIndex_Call) RunAndReturn(run func() math.U64) *Deposit_GetIndex_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *Deposit) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deposit_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type Deposit_MarshalSSZ_Call struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *Deposit_Expecter) MarshalSSZ() *Deposit_MarshalSSZ_Call {
	return &Deposit_MarshalSSZ_Call{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *Deposit_MarshalSSZ_Call) Run(run func()) *Deposit_MarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Deposit_MarshalSSZ_Call) Return(_a0 []byte, _a1 error) *Deposit_MarshalSSZ_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Deposit_MarshalSSZ_Call) RunAndReturn(run func() ([]byte, error)) *Deposit_MarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: _a0
func (_m *Deposit) UnmarshalSSZ(_a0 []byte) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deposit_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type Deposit_UnmarshalSSZ_Call struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - _a0 []byte
func (_e *Deposit_Expecter) UnmarshalSSZ(_a0 interface{}) *Deposit_UnmarshalSSZ_Call {
	return &Deposit_UnmarshalSSZ_Call{Call: _e.mock.On("UnmarshalSSZ", _a0)}
}

func (_c *Deposit_UnmarshalSSZ_Call) Run(run func(_a0 []byte)) *Deposit_UnmarshalSSZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *Deposit_UnmarshalSSZ_Call) Return(_a0 error) *Deposit_UnmarshalSSZ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_UnmarshalSSZ_Call) RunAndReturn(run func([]byte) error) *Deposit_UnmarshalSSZ_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeposit creates a new instance of Deposit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeposit(t interface {
	mock.TestingT
	Cleanup(func())
}) *Deposit {
	mock := &Deposit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
