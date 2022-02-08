// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sankethkini/ConcurrencyInGo/config (interfaces: ItaxRates)

// Package config is a generated GoMock package.
package config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockItaxRates is a mock of ItaxRates interface.
type MockItaxRates struct {
	ctrl     *gomock.Controller
	recorder *MockItaxRatesMockRecorder
}

// MockItaxRatesMockRecorder is the mock recorder for MockItaxRates.
type MockItaxRatesMockRecorder struct {
	mock *MockItaxRates
}

// NewMockItaxRates creates a new mock instance.
func NewMockItaxRates(ctrl *gomock.Controller) *MockItaxRates {
	mock := &MockItaxRates{ctrl: ctrl}
	mock.recorder = &MockItaxRatesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItaxRates) EXPECT() *MockItaxRatesMockRecorder {
	return m.recorder
}

// GetTaxRates mocks base method.
func (m *MockItaxRates) GetTaxRates() TaxRates {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaxRates")
	ret0, _ := ret[0].(TaxRates)
	return ret0
}

// GetTaxRates indicates an expected call of GetTaxRates.
func (mr *MockItaxRatesMockRecorder) GetTaxRates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaxRates", reflect.TypeOf((*MockItaxRates)(nil).GetTaxRates))
}
