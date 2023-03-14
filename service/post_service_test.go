package service_test

import (
	"babalaas/stella-artois/model"
	"babalaas/stella-artois/model/mocks"
	"babalaas/stella-artois/service"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchByUserID(t *testing.T) {
	mockPostRepo := new(mocks.PostRepository)
	usedPostID := uuid.New()
	notusedPostID := uuid.New()

	t.Run("success", func(t *testing.T) {

		mockPost := model.Post{
			ID:            usedPostID,
			UserProfileID: uuid.New(),
			CollectionID:  uuid.Nil,
			Caption:       "Test Caption",
			DateCreated:   time.Time{},
			Image:         "google.com",
			Image2:        "google.com",
			ReactionCount: 0,
			InCollection:  false,
		}

		mockPostRepo.On("GetByID", mock.Anything, usedPostID).Return(mockPost, nil).Once()

		service := service.NewPostService(mockPostRepo)

		returnedPost, err := service.GetByID(context.Background(), usedPostID)

		assert.NoError(t, err)
		assert.NotNil(t, returnedPost)
		assert.Equal(t, "Test Caption", returnedPost.Caption)
		mockPostRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		emptyPost := model.Post{}

		mockPostRepo.On("GetByID", mock.Anything, notusedPostID).Return(emptyPost, errors.New("Post with id not found.")).Once()

		service := service.NewPostService(mockPostRepo)

		returnedPost, err := service.GetByID(context.Background(), notusedPostID)

		assert.Error(t, err)
		assert.Equal(t, emptyPost, returnedPost)
		assert.Equal(t, returnedPost.ID, uuid.Nil)

		mockPostRepo.AssertExpectations(t)
	})
}
