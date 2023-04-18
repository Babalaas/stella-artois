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

func TestReactToPost(t *testing.T) {
	mockReactionRepo := new(mocks.ReactionRepository)

	serviceConfig := service.RSConfig{
		ReactionRepo: mockReactionRepo,
	}

	userProfileID := uuid.New()
	postID := uuid.New()
	reactionID := 1
	t.Run("success", func(t *testing.T) {
		mockPostReaction := model.PostReaction{
			// ID:            uuid.New(),
			UserProfileID: userProfileID,
			PostID:        postID,
			// DateCreated:   time.Time{},
			ReactionID: reactionID,
		}

		expected := model.PostReaction{
			ID:            uuid.New(),
			UserProfileID: userProfileID,
			PostID:        postID,
			DateCreated:   time.Now(),
			ReactionID:    reactionID,
		}

		mockReactionRepo.On("Create", mock.Anything, &mockPostReaction).Return(expected, nil).Once()

		service := service.NewReactionService(&serviceConfig)

		returnedReaction, err := service.ReactToPost(context.Background(), &mockPostReaction)

		assert.NoError(t, err)
		assert.Equal(t, expected.UserProfileID, returnedReaction.UserProfileID)
		assert.Equal(t, expected.PostID, returnedReaction.PostID)
		assert.NotNil(t, returnedReaction.DateCreated)
		assert.Equal(t, expected.ReactionID, returnedReaction.ReactionID)
		mockReactionRepo.AssertExpectations(t)
	})
}
