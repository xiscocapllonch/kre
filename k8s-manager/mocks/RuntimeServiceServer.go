// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import runtimepb "gitlab.com/konstellation/kre/k8s-manager/proto/runtimepb"

// RuntimeServiceServer is an autogenerated mock type for the RuntimeServiceServer type
type RuntimeServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *RuntimeServiceServer) Create(_a0 context.Context, _a1 *runtimepb.Request) (*runtimepb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *runtimepb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *runtimepb.Request) *runtimepb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtimepb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *runtimepb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuntimeStatus provides a mock function with given fields: _a0, _a1
func (_m *RuntimeServiceServer) RuntimeStatus(_a0 context.Context, _a1 *runtimepb.Request) (*runtimepb.RuntimeStatusResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *runtimepb.RuntimeStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *runtimepb.Request) *runtimepb.RuntimeStatusResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtimepb.RuntimeStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *runtimepb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
