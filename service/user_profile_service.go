package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"
)

// UserProfileService definition
type UserProfileService struct {
	UserProfileRepository model.UserProfileRepository
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
		return *&model.UserProfile{}, resErr
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
