// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	deposit "github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	mock "github.com/stretchr/testify/mock"

	types "github.com/berachain/beacon-kit/mod/async/pkg/types"
)

// BlockEvent is an autogenerated mock type for the BlockEvent type
type BlockEvent[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	mock.Mock
}

type BlockEvent_Expecter[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	mock *mock.Mock
}

func (_m *BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) EXPECT() *BlockEvent_Expecter[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	return &BlockEvent_Expecter[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]{mock: &_m.Mock}
}

// Data provides a mock function with given fields:
func (_m *BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Data() BeaconBlockT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 BeaconBlockT
	if rf, ok := ret.Get(0).(func() BeaconBlockT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BeaconBlockT)
	}

	return r0
}

// BlockEvent_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type BlockEvent_Data_Call[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Data() *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	return &BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]{Call: _e.mock.On("Data")}
}

func (_c *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Run(run func()) *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Return(_a0 BeaconBlockT) *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) RunAndReturn(run func() BeaconBlockT) *BlockEvent_Data_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(run)
	return _c
}

// Is provides a mock function with given fields: _a0
func (_m *BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Is(_a0 types.EventID) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Is")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.EventID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BlockEvent_Is_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Is'
type BlockEvent_Is_Call[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	*mock.Call
}

// Is is a helper method to define mock.On call
//   - _a0 types.EventID
func (_e *BlockEvent_Expecter[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Is(_a0 interface{}) *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	return &BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]{Call: _e.mock.On("Is", _a0)}
}

func (_c *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Run(run func(_a0 types.EventID)) *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.EventID))
	})
	return _c
}

func (_c *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Return(_a0 bool) *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) RunAndReturn(run func(types.EventID) bool) *BlockEvent_Is_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Type() types.EventID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 types.EventID
	if rf, ok := ret.Get(0).(func() types.EventID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.EventID)
	}

	return r0
}

// BlockEvent_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type BlockEvent_Type_Call[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *BlockEvent_Expecter[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Type() *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	return &BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]{Call: _e.mock.On("Type")}
}

func (_c *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Run(run func()) *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) Return(_a0 types.EventID) *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]) RunAndReturn(run func() types.EventID) *BlockEvent_Type_Call[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	_c.Call.Return(run)
	return _c
}

// NewBlockEvent creates a new instance of BlockEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockEvent[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT deposit.BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT] {
	mock := &BlockEvent[DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}