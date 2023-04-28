package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userProfileRepository struct {
	DB *gorm.DB
}

// SearchyByDisplayName implements model.UserProfileRepository
func (repo *userProfileRepository) SearchyByDisplayName(ctx context.Context, displayName string) ([]model.UserProfile, error) {
	var userProfiles []model.UserProfile
	query := "%" + displayName + "%"
	err := repo.DB.Where("display_name LIKE ?", query).Find(&userProfiles).Error
	if err != nil {
		return nil, err
	}
	return userProfiles, nil

}

// FindByID implements model.UserProfileRepository
func (repo *userProfileRepository) FindByID(ctx context.Context, userProfileID uuid.UUID) (model.UserProfile, error) {
	var resUserProfile model.UserProfile
	if resErr := repo.DB.Where("id = ?", userProfileID).First(&resUserProfile).Error; resErr != nil {
		log.Panic("User with id not found.")
		return resUserProfile, resErr
	}
	return resUserProfile, nil
}

// Create implements model.UserProfileRepository
func (repo *userProfileRepository) Create(ctx context.Context, userProfile *model.UserProfile) (model.UserProfile, error) {
	result := repo.DB.Create(&userProfile)

	if result.Error != nil {
		log.Panic("Could not create new User Profile.")
		return *userProfile, result.Error
	}

	return *userProfile, nil
}

// FindByDisplay name finds the first user_profile based on the passed display name
func (repo *userProfileRepository) FindByDisplayName(ctx context.Context, displayName string) (model.UserProfile, error) {
	var resUserProfile model.UserProfile
	if resErr := repo.DB.Where("display_name = ?", displayName).First(&resUserProfile).Error; resErr != nil {
		log.Panic("User with display name not found.")
		return resUserProfile, resErr
	}
	return resUserProfile, nil
}

// NewUserProfileRepository creates a new PostRepository with the server's database instance
func NewUserProfileRepository(db *gorm.DB) model.UserProfileRepository {
	return &userProfileRepository{
		DB: db,
	}
}
