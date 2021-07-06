// Code generated by mockery 2.7.5. DO NOT EDIT.

package rsk

import (
	rsk "github.com/rsksmart/rosetta-rsk/rsk"
	mock "github.com/stretchr/testify/mock"
)

// TransactionEncoder is an autogenerated mock type for the TransactionEncoder type
type TransactionEncoder struct {
	mock.Mock
}

// DecodeTransaction provides a mock function with given fields: _a0
func (_m *TransactionEncoder) DecodeTransaction(_a0 []byte) (*rsk.RlpTransactionParameters, error) {
	ret := _m.Called(_a0)

	var r0 *rsk.RlpTransactionParameters
	if rf, ok := ret.Get(0).(func([]byte) *rsk.RlpTransactionParameters); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsk.RlpTransactionParameters)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeRawTransaction provides a mock function with given fields: rlpTransactionParameters
func (_m *TransactionEncoder) EncodeRawTransaction(rlpTransactionParameters *rsk.RlpTransactionParameters) ([]byte, error) {
	ret := _m.Called(rlpTransactionParameters)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*rsk.RlpTransactionParameters) []byte); ok {
		r0 = rf(rlpTransactionParameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rsk.RlpTransactionParameters) error); ok {
		r1 = rf(rlpTransactionParameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeTransaction provides a mock function with given fields: rlpTransactionParameters
func (_m *TransactionEncoder) EncodeTransaction(rlpTransactionParameters *rsk.RlpTransactionParameters) ([]byte, error) {
	ret := _m.Called(rlpTransactionParameters)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*rsk.RlpTransactionParameters) []byte); ok {
		r0 = rf(rlpTransactionParameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rsk.RlpTransactionParameters) error); ok {
		r1 = rf(rlpTransactionParameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
