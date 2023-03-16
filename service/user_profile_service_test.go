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

func TestRegister(t *testing.T) {
	mockUserProfileRepo := new(mocks.UserProfileRepository)

	serviceConfig := service.UPSConfig{
		UserProfileRepository: mockUserProfileRepo,
	}
	t.Run("success", func(t *testing.T) {
		mockUserProfile := model.UserProfile{
			ID:          uuid.New(),
			DisplayName: "Dr. Brain May",
			FirstName:   "Brian",
			LastName:    "May",
			Email:       "brian.may@gmail.com",
			Phone:       "8141455656",
			Birthdate:   time.Time{},
			Password:    "queen",
			ProfilePic:  "google.com",
		}

		mockUserProfileRepo.On("Create", mock.Anything, &mockUserProfile).Return(mockUserProfile.ID, nil).Once()

		service := service.NewUserProfileService(&serviceConfig)

		returnedId, err := service.Register(context.Background(), &mockUserProfile)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, returnedId)
		mockUserProfileRepo.AssertExpectations(t)
	})
}
