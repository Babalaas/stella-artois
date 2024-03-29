// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CollectionService is an autogenerated mock type for the CollectionService type
type CollectionService struct {
	mock.Mock
}

// AddPostToCollection provides a mock function with given fields: ctx, postID, collectionID
func (_m *CollectionService) AddPostToCollection(ctx context.Context, postID uuid.UUID, collectionID uuid.UUID) error {
	ret := _m.Called(ctx, postID, collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, postID, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEmptyCollection provides a mock function with given fields: ctx, collection
func (_m *CollectionService) CreateEmptyCollection(ctx context.Context, collection model.Collection) error {
	ret := _m.Called(ctx, collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Collection) error); ok {
		r0 = rf(ctx, collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CollectionService) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserCollections provides a mock function with given fields: ctx, userProfileID
func (_m *CollectionService) GetUserCollections(ctx context.Context, userProfileID uuid.UUID) ([]model.Collection, error) {
	ret := _m.Called(ctx, userProfileID)

	var r0 []model.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.Collection, error)); ok {
		return rf(ctx, userProfileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.Collection); ok {
		r0 = rf(ctx, userProfileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userProfileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCollection provides a mock function with given fields: ctx, collection
func (_m *CollectionService) UpdateCollection(ctx context.Context, collection model.Collection) error {
	ret := _m.Called(ctx, collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Collection) error); ok {
		r0 = rf(ctx, collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCollectionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollectionService creates a new instance of CollectionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollectionService(t mockConstructorTestingTNewCollectionService) *CollectionService {
	mock := &CollectionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
