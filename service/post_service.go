package service

import (
	model "babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
)

type postService struct {
	PostRepository model.PostRepository
}

// UploadPost implements model.PostService
func (*postService) UploadPost(ctx context.Context, userProfileID uuid.UUID, caption string, image string) error {
	panic("unimplemented")
}

// Get implements models.PostService
func (service *postService) GetByID(ctx context.Context, uid uuid.UUID) (post model.Post, err error) {
	resPost, resErr := service.PostRepository.GetByID(ctx, uid)
	return resPost, resErr
}

func (service *postService) AddToCollection(ctx context.Context, post *model.Post) (err error) {
	panic("unimplemented")
}

// NewPostService creates a Post Service with a PostRepository attribute
func NewPostService(postRepo model.PostRepository) model.PostService {
	return &postService{
		PostRepository: postRepo,
	}
}
