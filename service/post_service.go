package service

import (
	"babalaas/stella-artois/models"
	"context"

	"github.com/google/uuid"
)

type postService struct {
	PostRepository models.PostRepository
}

// Get implements models.PostService
func (service *postService) GetById(ctx context.Context, uid uuid.UUID) (post models.Post, err error) {
	resPost, resErr := service.PostRepository.GetById(ctx, uid)
	return resPost, resErr
}

func (service *postService) AddToCollection(ctx context.Context, post *models.Post) (err error) {
	panic("unimplemented")
}

// can replace the parameter list with a PostServiceConfig parameter object
func NewPostService(postRepo *models.PostRepository) models.PostService {
	return &postService{
		PostRepository: *postRepo,
	}
}
