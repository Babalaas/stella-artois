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

func (repo *postRepository) Create(ctx context.Context, post *model.Post) (err error) {
	panic("unimplemented")
}

func (repo *postRepository) Delete(ctx context.Context, uid uuid.UUID) (err error) {
	panic("unimplemented")
}

func (repo *postRepository) GetAll(ctx context.Context) (posts []model.Post, err error) {
	panic("unimplemented")
}

func (repo *postRepository) GetByID(ctx context.Context, uid uuid.UUID) (post model.Post, err error) {
	var resPost model.Post
	if resErr := repo.DB.Where("id = ?", uid).First(&resPost).Error; resErr != nil {
		log.Panic("Post with id not found.")
		return resPost, resErr
	}
	return resPost, nil
}

func (repo *postRepository) Update(ctx context.Context, post *model.Post) (err error) {
	panic("unimplemented")
}

// NewPostRepository creates a new PostRepository with the server's database instance
func NewPostRepository(db *gorm.DB) model.PostRepository {
	return &postRepository{
		DB: db,
	}
}
