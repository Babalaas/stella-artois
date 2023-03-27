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

		actual := model.UserProfile{
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

		mockUserProfileRepo.On("Create", mock.Anything, &mockUserProfile).Return(actual, nil).Once()

		service := service.NewUserProfileService(&serviceConfig)

		returnedUserProfile, err := service.Register(context.Background(), &mockUserProfile)

		assert.NoError(t, err)
		assert.Equal(t, actual.DisplayName, returnedUserProfile.DisplayName)
		mockUserProfileRepo.AssertExpectations(t)
	})
}

func TestLogIn(t *testing.T) {
	mockUserProfileRepo := new(mocks.UserProfileRepository)

	serviceConfig := service.UPSConfig{
		UserProfileRepository: mockUserProfileRepo,
	}
	t.Run("success", func(t *testing.T) {
		mockLogInRequest := model.UserProfile{
			DisplayName: "Dr. Brain May",
			Password:    "queen",
		}

		mockUserProfileRepo.On("FindByDisplayName", mock.Anything, mockLogInRequest.DisplayName).Return(mockLogInRequest, nil).Once()

		service := service.NewUserProfileService(&serviceConfig)

		returnUser, err := service.LogIn(context.Background(), &mockLogInRequest)

		assert.NoError(t, err)
		assert.Equal(t, mockLogInRequest.DisplayName, returnUser.DisplayName)
		mockUserProfileRepo.AssertExpectations(t)
	})
}
