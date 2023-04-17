package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

// UserProfileService definition
type UserProfileService struct {
	UserProfileRepository model.UserProfileRepository
}

// GetDisplayName implements model.UserProfileService
func (service *UserProfileService) GetDisplayName(ctx context.Context, userProfileID uuid.UUID) (string, error) {
	userProfile, err := service.UserProfileRepository.FindByID(ctx, userProfileID)

	if err != nil {
		log.Fatal("Could not get user")
		return "", err
	}

	return userProfile.DisplayName, err

}

// UPSConfig defines the dependencies for a UserProfileService
type UPSConfig struct {
	UserProfileRepository model.UserProfileRepository
}

// Register implements model.UserProfileService
func (service *UserProfileService) Register(ctx context.Context, userProfile *model.UserProfile) (model.UserProfile, error) {
	_, resErr := service.UserProfileRepository.Create(ctx, userProfile)

	if resErr != nil {
		log.Panic("UserProfileService could not create new User Profile.")
		return *userProfile, resErr
	}

	return *userProfile, nil
}

// LogIn implements model.UserProfileService
func (service *UserProfileService) LogIn(ctx context.Context, userProfile *model.UserProfile) (model.UserProfile, error) {
	fetchedUserProfile, err := service.UserProfileRepository.FindByDisplayName(ctx, userProfile.DisplayName)

	if err != nil {
		log.Panic("UserProfileService could not fetch User Profile.")
		return fetchedUserProfile, err
	}

	// compare passwords
	doesMatch, err := comparePasswords(fetchedUserProfile.Password, userProfile.Password)

	if err != nil {
		log.Panic("UserProfileService: error processing password.")
	}

	if !doesMatch {
		log.Panic("UserProfileService: passwords do not match.")
	}

	return fetchedUserProfile, nil
}

// NewUserProfileService is a factory function for initialization a UserProfileService with its repository layer dependencies
func NewUserProfileService(config *UPSConfig) model.UserProfileService {
	return &UserProfileService{
		UserProfileRepository: config.UserProfileRepository,
	}
}
