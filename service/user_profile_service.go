package service

import (
	"babalaas/stella-artois/model"
	"context"
)

type UserProfileService struct {
	UserProfileRepository model.UserProfileRepository
}

type UPSConfig struct {
	UserProfileRepository model.UserProfileRepository
}

// Register implements model.UserProfileService
func (service *UserProfileService) Register(ctx context.Context, userProfile *model.UserProfile) (err error) {
	if resErr := service.UserProfileRepository.Create(ctx, userProfile); err != nil {
		return resErr
	}
	return nil
}

// NewPostService is a factory function for initialization a UserProfileService with its repository layer dependencies
func NewUserProfileService(config *UPSConfig) model.UserProfileService {
	return &UserProfileService{
		UserProfileRepository: config.UserProfileRepository,
	}
}
