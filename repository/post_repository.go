package repository

import (
	"babalaas/stella-artois/db"
	"babalaas/stella-artois/models"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func (repo *postRepository) Create(ctx context.Context, post *models.Post) (err error) {
	panic("unimplemented")
}

func (repo *postRepository) Delete(ctx context.Context, uid uuid.UUID) (err error) {
	panic("unimplemented")
}

func (repo *postRepository) GetAll(ctx context.Context) (posts []models.Post, err error) {
	panic("unimplemented")
}

func (repo *postRepository) GetById(ctx context.Context, uid uuid.UUID) (post models.Post, err error) {
	var resPost models.Post
	if resErr := repo.DB.Where("id = ?", uid).First(&resPost).Error; err != nil {
		log.Panic("Post with id not found.")
		return resPost, resErr
	}
	return resPost, nil
}

func (repo *postRepository) Update(ctx context.Context, post *models.Post) (err error) {
	panic("unimplemented")
}

func NewPostRepository() models.PostRepository {
	return &postRepository{
		DB: db.GetInstance(),
	}
}
