package service

import (
	"babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
)

type postCommentService struct {
	postCommentRepo model.CommentRepository
}

// GetAll implements model.CommentService
func (service *postCommentService) GetAll(ctx context.Context, postID uuid.UUID) ([]model.PostComment, error) {
	comments, err := service.postCommentRepo.GetRecent(ctx, postID, -1)
	return comments, err
}

// CSConfig is the parameter object for creating new CommentServices
type CSConfig struct {
	CommentRepo model.CommentRepository
}

// Create implements model.PostCommentService
func (service *postCommentService) Create(ctx context.Context, comment *model.PostComment) (model.PostComment, error) {
	newComment, err := service.postCommentRepo.Create(ctx, comment)
	return newComment, err
}

// Delete implements model.PostCommentService
func (service *postCommentService) Delete(ctx context.Context, comment *model.PostComment) error {
	err := service.postCommentRepo.Delete(ctx, comment)
	return err
}

// NewCommentService is the factory function for the CommentService struct
func NewCommentService(config *CSConfig) model.CommentService {
	return &postCommentService{
		postCommentRepo: config.CommentRepo,
	}
}
