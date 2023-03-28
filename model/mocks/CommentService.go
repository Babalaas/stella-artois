// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, comment
func (_m *CommentService) Create(ctx context.Context, comment *model.PostComment) (model.PostComment, error) {
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

// Delete provides a mock function with given fields: ctx, comment
func (_m *CommentService) Delete(ctx context.Context, comment *model.PostComment) error {
	ret := _m.Called(ctx, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.PostComment) error); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCommentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentService(t mockConstructorTestingTNewCommentService) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
