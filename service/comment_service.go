package service

import (
	"babalaas/stella-artois/model"
	"context"
)

type postCommentService struct {
	postCommentRepo model.CommentRepository
}

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

func NewCommentService(config *CSConfig) model.CommentService {
	return &postCommentService{
		postCommentRepo: config.CommentRepo,
	}
}
