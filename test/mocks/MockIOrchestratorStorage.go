// Code generated by mockery v2.46.2. DO NOT EDIT.

//go:build !production

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
	common "github.com/thirdweb-dev/indexer/internal/common"

	storage "github.com/thirdweb-dev/indexer/internal/storage"
)

// MockIOrchestratorStorage is an autogenerated mock type for the IOrchestratorStorage type
type MockIOrchestratorStorage struct {
	mock.Mock
}

type MockIOrchestratorStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIOrchestratorStorage) EXPECT() *MockIOrchestratorStorage_Expecter {
	return &MockIOrchestratorStorage_Expecter{mock: &_m.Mock}
}

// DeleteBlockFailures provides a mock function with given fields: failures
func (_m *MockIOrchestratorStorage) DeleteBlockFailures(failures []common.BlockFailure) error {
	ret := _m.Called(failures)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlockFailures")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]common.BlockFailure) error); ok {
		r0 = rf(failures)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOrchestratorStorage_DeleteBlockFailures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBlockFailures'
type MockIOrchestratorStorage_DeleteBlockFailures_Call struct {
	*mock.Call
}

// DeleteBlockFailures is a helper method to define mock.On call
//   - failures []common.BlockFailure
func (_e *MockIOrchestratorStorage_Expecter) DeleteBlockFailures(failures interface{}) *MockIOrchestratorStorage_DeleteBlockFailures_Call {
	return &MockIOrchestratorStorage_DeleteBlockFailures_Call{Call: _e.mock.On("DeleteBlockFailures", failures)}
}

func (_c *MockIOrchestratorStorage_DeleteBlockFailures_Call) Run(run func(failures []common.BlockFailure)) *MockIOrchestratorStorage_DeleteBlockFailures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]common.BlockFailure))
	})
	return _c
}

func (_c *MockIOrchestratorStorage_DeleteBlockFailures_Call) Return(_a0 error) *MockIOrchestratorStorage_DeleteBlockFailures_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOrchestratorStorage_DeleteBlockFailures_Call) RunAndReturn(run func([]common.BlockFailure) error) *MockIOrchestratorStorage_DeleteBlockFailures_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockFailures provides a mock function with given fields: qf
func (_m *MockIOrchestratorStorage) GetBlockFailures(qf storage.QueryFilter) ([]common.BlockFailure, error) {
	ret := _m.Called(qf)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockFailures")
	}

	var r0 []common.BlockFailure
	var r1 error
	if rf, ok := ret.Get(0).(func(storage.QueryFilter) ([]common.BlockFailure, error)); ok {
		return rf(qf)
	}
	if rf, ok := ret.Get(0).(func(storage.QueryFilter) []common.BlockFailure); ok {
		r0 = rf(qf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BlockFailure)
		}
	}

	if rf, ok := ret.Get(1).(func(storage.QueryFilter) error); ok {
		r1 = rf(qf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIOrchestratorStorage_GetBlockFailures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockFailures'
type MockIOrchestratorStorage_GetBlockFailures_Call struct {
	*mock.Call
}

// GetBlockFailures is a helper method to define mock.On call
//   - qf storage.QueryFilter
func (_e *MockIOrchestratorStorage_Expecter) GetBlockFailures(qf interface{}) *MockIOrchestratorStorage_GetBlockFailures_Call {
	return &MockIOrchestratorStorage_GetBlockFailures_Call{Call: _e.mock.On("GetBlockFailures", qf)}
}

func (_c *MockIOrchestratorStorage_GetBlockFailures_Call) Run(run func(qf storage.QueryFilter)) *MockIOrchestratorStorage_GetBlockFailures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(storage.QueryFilter))
	})
	return _c
}

func (_c *MockIOrchestratorStorage_GetBlockFailures_Call) Return(_a0 []common.BlockFailure, _a1 error) *MockIOrchestratorStorage_GetBlockFailures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIOrchestratorStorage_GetBlockFailures_Call) RunAndReturn(run func(storage.QueryFilter) ([]common.BlockFailure, error)) *MockIOrchestratorStorage_GetBlockFailures_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastReorgCheckedBlockNumber provides a mock function with given fields: chainId
func (_m *MockIOrchestratorStorage) GetLastReorgCheckedBlockNumber(chainId *big.Int) (*big.Int, error) {
	ret := _m.Called(chainId)

	if len(ret) == 0 {
		panic("no return value specified for GetLastReorgCheckedBlockNumber")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) (*big.Int, error)); ok {
		return rf(chainId)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(chainId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(chainId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastReorgCheckedBlockNumber'
type MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call struct {
	*mock.Call
}

// GetLastReorgCheckedBlockNumber is a helper method to define mock.On call
//   - chainId *big.Int
func (_e *MockIOrchestratorStorage_Expecter) GetLastReorgCheckedBlockNumber(chainId interface{}) *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call {
	return &MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call{Call: _e.mock.On("GetLastReorgCheckedBlockNumber", chainId)}
}

func (_c *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call) Run(run func(chainId *big.Int)) *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Int))
	})
	return _c
}

func (_c *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call) Return(_a0 *big.Int, _a1 error) *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call) RunAndReturn(run func(*big.Int) (*big.Int, error)) *MockIOrchestratorStorage_GetLastReorgCheckedBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// SetLastReorgCheckedBlockNumber provides a mock function with given fields: chainId, blockNumber
func (_m *MockIOrchestratorStorage) SetLastReorgCheckedBlockNumber(chainId *big.Int, blockNumber *big.Int) error {
	ret := _m.Called(chainId, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for SetLastReorgCheckedBlockNumber")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) error); ok {
		r0 = rf(chainId, blockNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLastReorgCheckedBlockNumber'
type MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call struct {
	*mock.Call
}

// SetLastReorgCheckedBlockNumber is a helper method to define mock.On call
//   - chainId *big.Int
//   - blockNumber *big.Int
func (_e *MockIOrchestratorStorage_Expecter) SetLastReorgCheckedBlockNumber(chainId interface{}, blockNumber interface{}) *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call {
	return &MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call{Call: _e.mock.On("SetLastReorgCheckedBlockNumber", chainId, blockNumber)}
}

func (_c *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call) Run(run func(chainId *big.Int, blockNumber *big.Int)) *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Int), args[1].(*big.Int))
	})
	return _c
}

func (_c *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call) Return(_a0 error) *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call) RunAndReturn(run func(*big.Int, *big.Int) error) *MockIOrchestratorStorage_SetLastReorgCheckedBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// StoreBlockFailures provides a mock function with given fields: failures
func (_m *MockIOrchestratorStorage) StoreBlockFailures(failures []common.BlockFailure) error {
	ret := _m.Called(failures)

	if len(ret) == 0 {
		panic("no return value specified for StoreBlockFailures")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]common.BlockFailure) error); ok {
		r0 = rf(failures)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOrchestratorStorage_StoreBlockFailures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreBlockFailures'
type MockIOrchestratorStorage_StoreBlockFailures_Call struct {
	*mock.Call
}

// StoreBlockFailures is a helper method to define mock.On call
//   - failures []common.BlockFailure
func (_e *MockIOrchestratorStorage_Expecter) StoreBlockFailures(failures interface{}) *MockIOrchestratorStorage_StoreBlockFailures_Call {
	return &MockIOrchestratorStorage_StoreBlockFailures_Call{Call: _e.mock.On("StoreBlockFailures", failures)}
}

func (_c *MockIOrchestratorStorage_StoreBlockFailures_Call) Run(run func(failures []common.BlockFailure)) *MockIOrchestratorStorage_StoreBlockFailures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]common.BlockFailure))
	})
	return _c
}

func (_c *MockIOrchestratorStorage_StoreBlockFailures_Call) Return(_a0 error) *MockIOrchestratorStorage_StoreBlockFailures_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOrchestratorStorage_StoreBlockFailures_Call) RunAndReturn(run func([]common.BlockFailure) error) *MockIOrchestratorStorage_StoreBlockFailures_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIOrchestratorStorage creates a new instance of MockIOrchestratorStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIOrchestratorStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIOrchestratorStorage {
	mock := &MockIOrchestratorStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}