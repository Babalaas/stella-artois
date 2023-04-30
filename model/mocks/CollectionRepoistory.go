// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	model "babalaas/stella-artois/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CollectionRepoistory is an autogenerated mock type for the CollectionRepoistory type
type CollectionRepoistory struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, collection
func (_m *CollectionRepoistory) Create(ctx context.Context, collection model.Collection) error {
	ret := _m.Called(ctx, collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Collection) error); ok {
		r0 = rf(ctx, collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCollectionRepoistory interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollectionRepoistory creates a new instance of CollectionRepoistory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollectionRepoistory(t mockConstructorTestingTNewCollectionRepoistory) *CollectionRepoistory {
	mock := &CollectionRepoistory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
