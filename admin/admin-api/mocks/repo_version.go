// Code generated by MockGen. DO NOT EDIT.
// Source: version.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/konstellation-io/kre/admin/admin-api/domain/entity"
	reflect "reflect"
)

// MockVersionRepo is a mock of VersionRepo interface
type MockVersionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockVersionRepoMockRecorder
}

// MockVersionRepoMockRecorder is the mock recorder for MockVersionRepo
type MockVersionRepoMockRecorder struct {
	mock *MockVersionRepo
}

// NewMockVersionRepo creates a new mock instance
func NewMockVersionRepo(ctrl *gomock.Controller) *MockVersionRepo {
	mock := &MockVersionRepo{ctrl: ctrl}
	mock.recorder = &MockVersionRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVersionRepo) EXPECT() *MockVersionRepoMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockVersionRepo) Create(userID string, version *entity.Version) (*entity.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, version)
	ret0, _ := ret[0].(*entity.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVersionRepoMockRecorder) Create(userID, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVersionRepo)(nil).Create), userID, version)
}

// GetByID mocks base method
func (m *MockVersionRepo) GetByID(id string) (*entity.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockVersionRepoMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVersionRepo)(nil).GetByID), id)
}

// GetByIDs mocks base method
func (m *MockVersionRepo) GetByIDs(ids []string) ([]*entity.Version, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ids)
	ret0, _ := ret[0].([]*entity.Version)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockVersionRepoMockRecorder) GetByIDs(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockVersionRepo)(nil).GetByIDs), ids)
}

// Update mocks base method
func (m *MockVersionRepo) Update(version *entity.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockVersionRepoMockRecorder) Update(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVersionRepo)(nil).Update), version)
}

// SetHasDoc mocks base method
func (m *MockVersionRepo) SetHasDoc(ctx context.Context, versionID string, hasDoc bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHasDoc", ctx, versionID, hasDoc)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHasDoc indicates an expected call of SetHasDoc
func (mr *MockVersionRepoMockRecorder) SetHasDoc(ctx, versionID, hasDoc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHasDoc", reflect.TypeOf((*MockVersionRepo)(nil).SetHasDoc), ctx, versionID, hasDoc)
}

// GetByRuntime mocks base method
func (m *MockVersionRepo) GetByRuntime(runtimeID string) ([]*entity.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRuntime", runtimeID)
	ret0, _ := ret[0].([]*entity.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRuntime indicates an expected call of GetByRuntime
func (mr *MockVersionRepoMockRecorder) GetByRuntime(runtimeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRuntime", reflect.TypeOf((*MockVersionRepo)(nil).GetByRuntime), runtimeID)
}
