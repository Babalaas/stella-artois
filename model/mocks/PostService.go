// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// AddToCollection provides a mock function with given fields: ctx, post
func (_m *PostService) AddToCollection(ctx context.Context, post *model.Post) error {
	ret := _m.Called(ctx, post)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Post) error); ok {
		r0 = rf(ctx, post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllByUserProfile provides a mock function with given fields: ctx, userProfileID
func (_m *PostService) GetAllByUserProfile(ctx context.Context, userProfileID uuid.UUID) ([]model.Post, error) {
	ret := _m.Called(ctx, userProfileID)

	var r0 []model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.Post, error)); ok {
		return rf(ctx, userProfileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.Post); ok {
		r0 = rf(ctx, userProfileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userProfileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, uid
func (_m *PostService) GetByID(ctx context.Context, uid uuid.UUID) (model.Post, error) {
	ret := _m.Called(ctx, uid)

	var r0 model.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (model.Post, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) model.Post); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Get(0).(model.Post)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPost provides a mock function with given fields: ctx, userProfileID, caption, image
func (_m *PostService) UploadPost(ctx context.Context, userProfileID uuid.UUID, caption string, image string) error {
	ret := _m.Called(ctx, userProfileID, caption, image)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, string) error); ok {
		r0 = rf(ctx, userProfileID, caption, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPostService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostService(t mockConstructorTestingTNewPostService) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
