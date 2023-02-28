package repository

import (
	"babalaas/stella-artois/db"
	"babalaas/stella-artois/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type userProfileRepository struct {
	DB *gorm.DB
}

// Create implements model.UserProfileRepository
func (repo *userProfileRepository) Create(ctx context.Context, userProfile *model.UserProfile) (err error) {
	if resErr := repo.DB.Create(&userProfile).Error; resErr != nil {
		log.Panic("Could not create new User Profile.")
		return resErr
	}
	return nil
}

// NewPostRepository creates a new PostRepository with the server's database instance
func NewUserProfileRepository() model.UserProfileRepository {
	return &userProfileRepository{
		DB: db.GetInstance(),
	}
}

// NewPostRepository creates a new PostRepository with the server's database instance
func NewTestUserProfileRepository(db *gorm.DB) model.UserProfileRepository {
	return &userProfileRepository{
		DB: db,
	}
}
