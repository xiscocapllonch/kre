// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	monitoringpb "github.com/konstellation-io/kre/runtime/runtime-api/proto/monitoringpb"
)

// MonitoringServiceClient is an autogenerated mock type for the MonitoringServiceClient type
type MonitoringServiceClient struct {
	mock.Mock
}

// GetMetrics provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringServiceClient) GetMetrics(ctx context.Context, in *monitoringpb.GetMetricsRequest, opts ...grpc.CallOption) (*monitoringpb.GetMetricsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *monitoringpb.GetMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *monitoringpb.GetMetricsRequest, ...grpc.CallOption) *monitoringpb.GetMetricsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitoringpb.GetMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *monitoringpb.GetMetricsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeLogs provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringServiceClient) NodeLogs(ctx context.Context, in *monitoringpb.NodeLogsRequest, opts ...grpc.CallOption) (monitoringpb.MonitoringService_NodeLogsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 monitoringpb.MonitoringService_NodeLogsClient
	if rf, ok := ret.Get(0).(func(context.Context, *monitoringpb.NodeLogsRequest, ...grpc.CallOption) monitoringpb.MonitoringService_NodeLogsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monitoringpb.MonitoringService_NodeLogsClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *monitoringpb.NodeLogsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchLogs provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringServiceClient) SearchLogs(ctx context.Context, in *monitoringpb.SearchLogsRequest, opts ...grpc.CallOption) (*monitoringpb.SearchLogsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *monitoringpb.SearchLogsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *monitoringpb.SearchLogsRequest, ...grpc.CallOption) *monitoringpb.SearchLogsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitoringpb.SearchLogsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *monitoringpb.SearchLogsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchNodeStatus provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringServiceClient) WatchNodeStatus(ctx context.Context, in *monitoringpb.NodeStatusRequest, opts ...grpc.CallOption) (monitoringpb.MonitoringService_WatchNodeStatusClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 monitoringpb.MonitoringService_WatchNodeStatusClient
	if rf, ok := ret.Get(0).(func(context.Context, *monitoringpb.NodeStatusRequest, ...grpc.CallOption) monitoringpb.MonitoringService_WatchNodeStatusClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monitoringpb.MonitoringService_WatchNodeStatusClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *monitoringpb.NodeStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
