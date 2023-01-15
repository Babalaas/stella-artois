package service

import (
	model "babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
)

type postService struct {
	PostRepository model.PostRepository
}

// Get implements models.PostService
func (service *postService) GetById(ctx context.Context, uid uuid.UUID) (post model.Post, err error) {
	resPost, resErr := service.PostRepository.GetById(ctx, uid)
	return resPost, resErr
}

func (service *postService) AddToCollection(ctx context.Context, post *model.Post) (err error) {
	panic("unimplemented")
}

// NewPostService creates a Post Service with a PostRepository attribute
func NewPostService(postRepo *model.PostRepository) model.PostService {
	return &postService{
		PostRepository: *postRepo,
	}
}
