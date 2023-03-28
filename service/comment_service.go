package service

import (
	"babalaas/stella-artois/model"
	"context"
)

type postCommentService struct {
	postCommentRepo model.CommentRepository
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
