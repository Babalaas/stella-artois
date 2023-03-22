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

// UPSConfig defines the dependencies for a UserProfileService
type UPSConfig struct {
	UserProfileRepository model.UserProfileRepository
}

// Register implements model.UserProfileService
func (service *UserProfileService) Register(ctx context.Context, userProfile *model.UserProfile) (uuid.UUID, error) {
	newID, resErr := service.UserProfileRepository.Create(ctx, userProfile)

	if resErr != nil {
		log.Panic("UserProfileService could not create new User Profile.")
		return uuid.Nil, resErr
	}

	return newID, nil
}

// LogIn implements model.UserProfileService
func (service *UserProfileService) LogIn(ctx context.Context, userProfile *model.UserProfile) error {
	fetchedUserProfile, err := service.UserProfileRepository.FindByDisplayName(ctx, userProfile.DisplayName)

	if err != nil {
		log.Panic("UserProfileService could not fetch User Profile.")
		return err
	}

	// compare passwords
	doesMatch, err := comparePasswords(fetchedUserProfile.Password, userProfile.Password)

	if err != nil {
		log.Panic("UserProfileService: error processing password.")
	}

	if !doesMatch {
		log.Panic("UserProfileService: passwords do not match.")
	}

	return nil
}

// NewUserProfileService is a factory function for initialization a UserProfileService with its repository layer dependencies
func NewUserProfileService(config *UPSConfig) model.UserProfileService {
	return &UserProfileService{
		UserProfileRepository: config.UserProfileRepository,
	}
}
