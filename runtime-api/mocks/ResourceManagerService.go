// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "gitlab.com/konstellation/konstellation-ce/kre/runtime-api/domain/entity"
import mock "github.com/stretchr/testify/mock"

// ResourceManagerService is an autogenerated mock type for the ResourceManagerService type
type ResourceManagerService struct {
	mock.Mock
}

// CreateEntrypoint provides a mock function with given fields: version
func (_m *ResourceManagerService) CreateEntrypoint(version *entity.Version) error {
	ret := _m.Called(version)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Version) error); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNode provides a mock function with given fields: version, node
func (_m *ResourceManagerService) CreateNode(version *entity.Version, node *entity.Node) error {
	ret := _m.Called(version, node)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Version, *entity.Node) error); ok {
		r0 = rf(version, node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateVersionConfig provides a mock function with given fields: version
func (_m *ResourceManagerService) CreateVersionConfig(version *entity.Version) (string, error) {
	ret := _m.Called(version)

	var r0 string
	if rf, ok := ret.Get(0).(func(*entity.Version) string); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Version) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishVersion provides a mock function with given fields: name
func (_m *ResourceManagerService) PublishVersion(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopVersion provides a mock function with given fields: name
func (_m *ResourceManagerService) StopVersion(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnpublishVersion provides a mock function with given fields: name
func (_m *ResourceManagerService) UnpublishVersion(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateVersionConfig provides a mock function with given fields: version
func (_m *ResourceManagerService) UpdateVersionConfig(version *entity.Version) error {
	ret := _m.Called(version)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Version) error); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchVersionNodeStatus provides a mock function with given fields: versionName, statusCh
func (_m *ResourceManagerService) WatchVersionNodeStatus(versionName string, statusCh chan<- *entity.VersionNodeStatus) chan struct{} {
	ret := _m.Called(versionName, statusCh)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(string, chan<- *entity.VersionNodeStatus) chan struct{}); ok {
		r0 = rf(versionName, statusCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	return r0
}
