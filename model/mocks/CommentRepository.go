// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CommentRepository is an autogenerated mock type for the CommentRepository type
type CommentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, comment
func (_m *CommentRepository) Create(ctx context.Context, comment *model.PostComment) (model.PostComment, error) {
	ret := _m.Called(ctx, comment)

	var r0 model.PostComment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.PostComment) (model.PostComment, error)); ok {
		return rf(ctx, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.PostComment) model.PostComment); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Get(0).(model.PostComment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.PostComment) error); ok {
		r1 = rf(ctx, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, commentID
func (_m *CommentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	ret := _m.Called(ctx, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRecent provides a mock function with given fields: ctx, postID, limit
func (_m *CommentRepository) GetRecent(ctx context.Context, postID uuid.UUID, limit int) ([]model.PostComment, error) {
	ret := _m.Called(ctx, postID, limit)

	var r0 []model.PostComment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) ([]model.PostComment, error)); ok {
		return rf(ctx, postID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) []model.PostComment); ok {
		r0 = rf(ctx, postID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PostComment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int) error); ok {
		r1 = rf(ctx, postID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentRepository creates a new instance of CommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentRepository(t mockConstructorTestingTNewCommentRepository) *CommentRepository {
	mock := &CommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
