package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

// GetAllByUserProfile implements model.PostRepository
func (*postRepository) GetAllByUserProfile(ctx context.Context, userProfileID uuid.UUID) ([]model.Post, error) {
	panic("unimplemented")
}

func (repo *postRepository) Create(ctx context.Context, post model.Post) error {
	err := repo.DB.Create(&post).Error
	return err
}

func (repo *postRepository) GetByID(ctx context.Context, uid uuid.UUID) (post model.Post, err error) {
	var resPost model.Post
	if resErr := repo.DB.Where("id = ?", uid).First(&resPost).Error; resErr != nil {
		log.Panic("Post with id not found.")
		return resPost, resErr
	}
	return resPost, nil
}

// NewPostRepository creates a new PostRepository with the server's database instance
func NewPostRepository(db *gorm.DB) model.PostRepository {
	return &postRepository{
		DB: db,
	}
}
