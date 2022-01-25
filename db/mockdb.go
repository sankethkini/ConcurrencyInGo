// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sankethkini/ConcurrencyInGo/db (interfaces: DBHelper)

// Package mock_db is a generated GoMock package.
package db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sankethkini/ConcurrencyInGo/model"
)

// MockDBHelper is a mock of DBHelper interface.
type MockDBHelper struct {
	ctrl     *gomock.Controller
	recorder *MockDBHelperMockRecorder
}

// MockDBHelperMockRecorder is the mock recorder for MockDBHelper.
type MockDBHelperMockRecorder struct {
	mock *MockDBHelper
}

// NewMockDBHelper creates a new mock instance.
func NewMockDBHelper(ctrl *gomock.Controller) *MockDBHelper {
	mock := &MockDBHelper{ctrl: ctrl}
	mock.recorder = &MockDBHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBHelper) EXPECT() *MockDBHelperMockRecorder {
	return m.recorder
}

// ReadDB mocks base method.
func (m *MockDBHelper) ReadDB() ([]model.BaseItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDB")
	ret0, _ := ret[0].([]model.BaseItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDB indicates an expected call of ReadDB.
func (mr *MockDBHelperMockRecorder) ReadDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDB", reflect.TypeOf((*MockDBHelper)(nil).ReadDB))
}
