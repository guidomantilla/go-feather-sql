// Code generated by MockGen. DO NOT EDIT.
// Source: jukebox-app/pkg/datasource (interfaces: RelationalDataSource)

// Package datasource is a generated GoMock package.
package datasource

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRelationalDataSource is a mock of RelationalDataSource interface.
type MockRelationalDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockRelationalDataSourceMockRecorder
}

// MockRelationalDataSourceMockRecorder is the mock recorder for MockRelationalDataSource.
type MockRelationalDataSourceMockRecorder struct {
	mock *MockRelationalDataSource
}

// NewMockRelationalDataSource creates a new mock instance.
func NewMockRelationalDataSource(ctrl *gomock.Controller) *MockRelationalDataSource {
	mock := &MockRelationalDataSource{ctrl: ctrl}
	mock.recorder = &MockRelationalDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelationalDataSource) EXPECT() *MockRelationalDataSourceMockRecorder {
	return m.recorder
}

// GetDatabase mocks base method.
func (m *MockRelationalDataSource) GetDatabase() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockRelationalDataSourceMockRecorder) GetDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockRelationalDataSource)(nil).GetDatabase))
}

// GetDriverName mocks base method.
func (m *MockRelationalDataSource) GetDriverName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDriverName indicates an expected call of GetDriverName.
func (mr *MockRelationalDataSourceMockRecorder) GetDriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverName", reflect.TypeOf((*MockRelationalDataSource)(nil).GetDriverName))
}
