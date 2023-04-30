// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// FriendshipService is an autogenerated mock type for the FriendshipService type
type FriendshipService struct {
	mock.Mock
}

// AcceptFriend provides a mock function with given fields: ctx, userProfileID, friendID
func (_m *FriendshipService) AcceptFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	ret := _m.Called(ctx, userProfileID, friendID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userProfileID, friendID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllFriends provides a mock function with given fields: ctx, userProfileID
func (_m *FriendshipService) GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]model.Friend, error) {
	ret := _m.Called(ctx, userProfileID)

	var r0 []model.Friend
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.Friend, error)); ok {
		return rf(ctx, userProfileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.Friend); ok {
		r0 = rf(ctx, userProfileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Friend)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userProfileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFriendRequests provides a mock function with given fields: ctx, userProfileID
func (_m *FriendshipService) GetFriendRequests(ctx context.Context, userProfileID uuid.UUID) ([]model.UserProfile, error) {
	ret := _m.Called(ctx, userProfileID)

	var r0 []model.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.UserProfile, error)); ok {
		return rf(ctx, userProfileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.UserProfile); ok {
		r0 = rf(ctx, userProfileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.UserProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userProfileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFriend provides a mock function with given fields: ctx, userProfileID, friendID
func (_m *FriendshipService) RemoveFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	ret := _m.Called(ctx, userProfileID, friendID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userProfileID, friendID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestFriend provides a mock function with given fields: ctx, userProfileID, friendID
func (_m *FriendshipService) RequestFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	ret := _m.Called(ctx, userProfileID, friendID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, userProfileID, friendID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFriendshipService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFriendshipService creates a new instance of FriendshipService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFriendshipService(t mockConstructorTestingTNewFriendshipService) *FriendshipService {
	mock := &FriendshipService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
