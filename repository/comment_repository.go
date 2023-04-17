package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

// Create implements model.PostCommentRepository
func (repo *commentRepository) Create(ctx context.Context, comment *model.PostComment) (model.PostComment, error) {
	result := repo.DB.Create(&comment)

	if result.Error != nil {
		log.Panic("Could not create new Post Comment.")
		return *comment, result.Error
	}

	return *comment, nil
}

// Delete implements model.PostCommentRepository
func (repo *commentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	var comment model.PostComment

	repo.DB.Where("id = ?", commentID).First(&comment)

	if comment.ID == uuid.Nil {
		log.Panic("Could not find comment with given ID")
	}

	result := repo.DB.Where("id = ?", commentID).Delete(comment)

	if result.Error != nil {
		log.Panic("Could not delete  Post Comment.")
		return result.Error
	}

	return nil
}

// NewCommentRepository is the factory function for creating CommentRepositories
func NewCommentRepository(db *gorm.DB) model.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}
