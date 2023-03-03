package repository

import (
	"babalaas/stella-artois/db"
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userProfileRepository struct {
	DB *gorm.DB
}

// Create implements model.UserProfileRepository
func (repo *userProfileRepository) Create(ctx context.Context, userProfile *model.UserProfile) (uuid.UUID, error) {
	result := repo.DB.Create(&userProfile)

	if result.Error != nil {
		log.Panic("Could not create new User Profile.")
		return uuid.Nil, result.Error
	}

	return userProfile.ID, nil
}

// NewUserProfileRepository creates a new PostRepository with the server's database instance
func NewUserProfileRepository() model.UserProfileRepository {
	return &userProfileRepository{
		DB: db.GetInstance(),
	}
}

// NewTestUserProfileRepository creates a new gorm.DB
// which is used with a sqlmock
func NewTestUserProfileRepository(db *gorm.DB) model.UserProfileRepository {
	return &userProfileRepository{
		DB: db,
	}
}
