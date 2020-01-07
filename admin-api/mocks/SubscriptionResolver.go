// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import gql "gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/gql"
import mock "github.com/stretchr/testify/mock"

// SubscriptionResolver is an autogenerated mock type for the SubscriptionResolver type
type SubscriptionResolver struct {
	mock.Mock
}

// RuntimeCreated provides a mock function with given fields: ctx
func (_m *SubscriptionResolver) RuntimeCreated(ctx context.Context) (<-chan *gql.Runtime, error) {
	ret := _m.Called(ctx)

	var r0 <-chan *gql.Runtime
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *gql.Runtime); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *gql.Runtime)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}