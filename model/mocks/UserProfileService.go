// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserProfileService is an autogenerated mock type for the UserProfileService type
type UserProfileService struct {
	mock.Mock
}

// LogIn provides a mock function with given fields: ctx, userProfile
func (_m *UserProfileService) LogIn(ctx context.Context, userProfile *model.UserProfile) error {
	ret := _m.Called(ctx, userProfile)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserProfile) error); ok {
		r0 = rf(ctx, userProfile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: ctx, userProfile
func (_m *UserProfileService) Register(ctx context.Context, userProfile *model.UserProfile) (uuid.UUID, error) {
	ret := _m.Called(ctx, userProfile)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserProfile) (uuid.UUID, error)); ok {
		return rf(ctx, userProfile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserProfile) uuid.UUID); ok {
		r0 = rf(ctx, userProfile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserProfile) error); ok {
		r1 = rf(ctx, userProfile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserProfileService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserProfileService creates a new instance of UserProfileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserProfileService(t mockConstructorTestingTNewUserProfileService) *UserProfileService {
	mock := &UserProfileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}