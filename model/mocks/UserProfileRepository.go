// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserProfileRepository is an autogenerated mock type for the UserProfileRepository type
type UserProfileRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, userProfile
func (_m *UserProfileRepository) Create(ctx context.Context, userProfile *model.UserProfile) (model.UserProfile, error) {
	ret := _m.Called(ctx, userProfile)

	var r0 model.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserProfile) (model.UserProfile, error)); ok {
		return rf(ctx, userProfile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserProfile) model.UserProfile); ok {
		r0 = rf(ctx, userProfile)
	} else {
		r0 = ret.Get(0).(model.UserProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserProfile) error); ok {
		r1 = rf(ctx, userProfile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByDisplayName provides a mock function with given fields: ctx, displayName
func (_m *UserProfileRepository) FindByDisplayName(ctx context.Context, displayName string) (model.UserProfile, error) {
	ret := _m.Called(ctx, displayName)

	var r0 model.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.UserProfile, error)); ok {
		return rf(ctx, displayName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.UserProfile); ok {
		r0 = rf(ctx, displayName)
	} else {
		r0 = ret.Get(0).(model.UserProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, displayName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, userProfileID
func (_m *UserProfileRepository) FindByID(ctx context.Context, userProfileID uuid.UUID) (model.UserProfile, error) {
	ret := _m.Called(ctx, userProfileID)

	var r0 model.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (model.UserProfile, error)); ok {
		return rf(ctx, userProfileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) model.UserProfile); ok {
		r0 = rf(ctx, userProfileID)
	} else {
		r0 = ret.Get(0).(model.UserProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userProfileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchyByDisplayName provides a mock function with given fields: ctx, displayName
func (_m *UserProfileRepository) SearchyByDisplayName(ctx context.Context, displayName string) ([]model.UserProfile, error) {
	ret := _m.Called(ctx, displayName)

	var r0 []model.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.UserProfile, error)); ok {
		return rf(ctx, displayName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.UserProfile); ok {
		r0 = rf(ctx, displayName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.UserProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, displayName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserProfileRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserProfileRepository creates a new instance of UserProfileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserProfileRepository(t mockConstructorTestingTNewUserProfileRepository) *UserProfileRepository {
	mock := &UserProfileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
