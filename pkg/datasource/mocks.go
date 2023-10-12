// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/datasource/types.go

// Package datasource is a generated GoMock package.
package datasource

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	sqlx "github.com/jmoiron/sqlx"
)

// MockDatasourceContext is a mock of DatasourceContext interface.
type MockDatasourceContext struct {
	ctrl     *gomock.Controller
	recorder *MockDatasourceContextMockRecorder
}

// MockDatasourceContextMockRecorder is the mock recorder for MockDatasourceContext.
type MockDatasourceContextMockRecorder struct {
	mock *MockDatasourceContext
}

// NewMockDatasourceContext creates a new mock instance.
func NewMockDatasourceContext(ctrl *gomock.Controller) *MockDatasourceContext {
	mock := &MockDatasourceContext{ctrl: ctrl}
	mock.recorder = &MockDatasourceContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasourceContext) EXPECT() *MockDatasourceContextMockRecorder {
	return m.recorder
}

// GetDriverName mocks base method.
func (m *MockDatasourceContext) GetDriverName() sql.DriverName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverName")
	ret0, _ := ret[0].(sql.DriverName)
	return ret0
}

// GetDriverName indicates an expected call of GetDriverName.
func (mr *MockDatasourceContextMockRecorder) GetDriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverName", reflect.TypeOf((*MockDatasourceContext)(nil).GetDriverName))
}

// GetParamHolder mocks base method.
func (m *MockDatasourceContext) GetParamHolder() sql.ParamHolder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParamHolder")
	ret0, _ := ret[0].(sql.ParamHolder)
	return ret0
}

// GetParamHolder indicates an expected call of GetParamHolder.
func (mr *MockDatasourceContextMockRecorder) GetParamHolder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParamHolder", reflect.TypeOf((*MockDatasourceContext)(nil).GetParamHolder))
}

// GetServer mocks base method.
func (m *MockDatasourceContext) GetServer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServer indicates an expected call of GetServer.
func (mr *MockDatasourceContextMockRecorder) GetServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockDatasourceContext)(nil).GetServer))
}

// GetService mocks base method.
func (m *MockDatasourceContext) GetService() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetService indicates an expected call of GetService.
func (mr *MockDatasourceContextMockRecorder) GetService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockDatasourceContext)(nil).GetService))
}

// GetUrl mocks base method.
func (m *MockDatasourceContext) GetUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockDatasourceContextMockRecorder) GetUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockDatasourceContext)(nil).GetUrl))
}

// MockDatasource is a mock of Datasource interface.
type MockDatasource struct {
	ctrl     *gomock.Controller
	recorder *MockDatasourceMockRecorder
}

// MockDatasourceMockRecorder is the mock recorder for MockDatasource.
type MockDatasourceMockRecorder struct {
	mock *MockDatasource
}

// NewMockDatasource creates a new mock instance.
func NewMockDatasource(ctrl *gomock.Controller) *MockDatasource {
	mock := &MockDatasource{ctrl: ctrl}
	mock.recorder = &MockDatasourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasource) EXPECT() *MockDatasourceMockRecorder {
	return m.recorder
}

// GetDatabase mocks base method.
func (m *MockDatasource) GetDatabase() (*sqlx.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(*sqlx.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockDatasourceMockRecorder) GetDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockDatasource)(nil).GetDatabase))
}

// MockTransactionHandler is a mock of TransactionHandler interface.
type MockTransactionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionHandlerMockRecorder
}

// MockTransactionHandlerMockRecorder is the mock recorder for MockTransactionHandler.
type MockTransactionHandlerMockRecorder struct {
	mock *MockTransactionHandler
}

// NewMockTransactionHandler creates a new mock instance.
func NewMockTransactionHandler(ctrl *gomock.Controller) *MockTransactionHandler {
	mock := &MockTransactionHandler{ctrl: ctrl}
	mock.recorder = &MockTransactionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionHandler) EXPECT() *MockTransactionHandlerMockRecorder {
	return m.recorder
}

// HandleTransaction mocks base method.
func (m *MockTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTransaction indicates an expected call of HandleTransaction.
func (mr *MockTransactionHandlerMockRecorder) HandleTransaction(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTransaction", reflect.TypeOf((*MockTransactionHandler)(nil).HandleTransaction), ctx, fn)
}
