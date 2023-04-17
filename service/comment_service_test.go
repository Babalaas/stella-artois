package service_test

import (
	"babalaas/stella-artois/model"
	"babalaas/stella-artois/model/mocks"
	"babalaas/stella-artois/service"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePostComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)

	serviceConfig := service.CSConfig{
		CommentRepo: mockCommentRepo,
	}
	t.Run("success", func(t *testing.T) {
		mockID := uuid.New()
		mockUserProfileID := uuid.New()
		mockPostID := uuid.New()

		mockPostComment := model.PostComment{
			UserProfileID: mockUserProfileID,
			PostID:        mockPostID,
			Content:       "This is a test post comment!",
		}

		expected := model.PostComment{
			ID:            mockID,
			UserProfileID: mockUserProfileID,
			PostID:        mockPostID,
			DateCreated:   time.Date(2023, time.March, 28, 14, 24, 50, 0, time.Local),
			Content:       "This is a test post comment!",
		}

		mockCommentRepo.On("Create", mock.Anything, &mockPostComment).Return(expected, nil).Once()

		service := service.NewCommentService(&serviceConfig)

		returnedPostComment, err := service.Create(context.Background(), &mockPostComment)

		assert.NoError(t, err)
		assert.Equal(t, expected.UserProfileID, returnedPostComment.UserProfileID)
		assert.NotNil(t, returnedPostComment.DateCreated)
		assert.Equal(t, expected.ID, returnedPostComment.ID)
		mockCommentRepo.AssertExpectations(t)
	})
}

func TestDeletePostComment(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepository)

	serviceConfig := service.CSConfig{
		CommentRepo: mockCommentRepo,
	}

	mockID := uuid.New()
	mockUserProfileID := uuid.New()
	mockPostID := uuid.New()

	service := service.NewCommentService(&serviceConfig)

	t.Run("success", func(t *testing.T) {
		mockPostComment := model.PostComment{
			ID:            mockID,
			UserProfileID: mockUserProfileID,
			PostID:        mockPostID,
			DateCreated:   time.Date(2023, time.March, 28, 14, 24, 50, 0, time.Local),
			Content:       "This is a test post comment!",
		}

		mockCommentRepo.On("Delete", mock.Anything, mockPostComment.ID).Return(nil).Once()
		err := service.Delete(context.Background(), mockPostComment.ID)

		assert.NoError(t, err)
		assert.Nil(t, err)
		mockCommentRepo.AssertExpectations(t)
	})

}
