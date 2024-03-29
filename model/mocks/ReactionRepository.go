// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ReactionRepository is an autogenerated mock type for the ReactionRepository type
type ReactionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, reaction
func (_m *ReactionRepository) Create(ctx context.Context, reaction *model.PostReaction) (model.PostReaction, error) {
	ret := _m.Called(ctx, reaction)

	var r0 model.PostReaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.PostReaction) (model.PostReaction, error)); ok {
		return rf(ctx, reaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.PostReaction) model.PostReaction); ok {
		r0 = rf(ctx, reaction)
	} else {
		r0 = ret.Get(0).(model.PostReaction)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.PostReaction) error); ok {
		r1 = rf(ctx, reaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewReactionRepository creates a new instance of ReactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReactionRepository(t mockConstructorTestingTNewReactionRepository) *ReactionRepository {
	mock := &ReactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
