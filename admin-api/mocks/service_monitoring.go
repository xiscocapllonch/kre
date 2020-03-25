// Code generated by MockGen. DO NOT EDIT.
// Source: monitoring.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "gitlab.com/konstellation/kre/admin-api/domain/entity"
	reflect "reflect"
)

// MockMonitoringService is a mock of MonitoringService interface
type MockMonitoringService struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringServiceMockRecorder
}

// MockMonitoringServiceMockRecorder is the mock recorder for MockMonitoringService
type MockMonitoringServiceMockRecorder struct {
	mock *MockMonitoringService
}

// NewMockMonitoringService creates a new mock instance
func NewMockMonitoringService(ctrl *gomock.Controller) *MockMonitoringService {
	mock := &MockMonitoringService{ctrl: ctrl}
	mock.recorder = &MockMonitoringServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitoringService) EXPECT() *MockMonitoringServiceMockRecorder {
	return m.recorder
}

// NodeLogs mocks base method
func (m *MockMonitoringService) NodeLogs(runtime *entity.Runtime, nodeID string, stopCh <-chan bool) (<-chan *entity.NodeLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeLogs", runtime, nodeID, stopCh)
	ret0, _ := ret[0].(<-chan *entity.NodeLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeLogs indicates an expected call of NodeLogs
func (mr *MockMonitoringServiceMockRecorder) NodeLogs(runtime, nodeID, stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeLogs", reflect.TypeOf((*MockMonitoringService)(nil).NodeLogs), runtime, nodeID, stopCh)
}

// VersionStatus mocks base method
func (m *MockMonitoringService) VersionStatus(runtime *entity.Runtime, versionName string, stopCh <-chan bool) (<-chan *entity.VersionNodeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionStatus", runtime, versionName, stopCh)
	ret0, _ := ret[0].(<-chan *entity.VersionNodeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VersionStatus indicates an expected call of VersionStatus
func (mr *MockMonitoringServiceMockRecorder) VersionStatus(runtime, versionName, stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionStatus", reflect.TypeOf((*MockMonitoringService)(nil).VersionStatus), runtime, versionName, stopCh)
}

// SearchLogs mocks base method
func (m *MockMonitoringService) SearchLogs(ctx context.Context, runtime *entity.Runtime, options entity.SearchLogsOptions) (entity.SearchLogsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLogs", ctx, runtime, options)
	ret0, _ := ret[0].(entity.SearchLogsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLogs indicates an expected call of SearchLogs
func (mr *MockMonitoringServiceMockRecorder) SearchLogs(ctx, runtime, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLogs", reflect.TypeOf((*MockMonitoringService)(nil).SearchLogs), ctx, runtime, options)
}
