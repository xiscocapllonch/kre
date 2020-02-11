// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/entity"
import gql "gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/gql"
import mock "github.com/stretchr/testify/mock"

// UserActivityResolver is an autogenerated mock type for the UserActivityResolver type
type UserActivityResolver struct {
	mock.Mock
}

// Date provides a mock function with given fields: ctx, obj
func (_m *UserActivityResolver) Date(ctx context.Context, obj *entity.UserActivity) (string, error) {
	ret := _m.Called(ctx, obj)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserActivity) string); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.UserActivity) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Type provides a mock function with given fields: ctx, obj
func (_m *UserActivityResolver) Type(ctx context.Context, obj *entity.UserActivity) (gql.UserActivityType, error) {
	ret := _m.Called(ctx, obj)

	var r0 gql.UserActivityType
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserActivity) gql.UserActivityType); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Get(0).(gql.UserActivityType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.UserActivity) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
