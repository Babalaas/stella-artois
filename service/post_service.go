package service

import (
	model "babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
)

type postService struct {
	bucketURL      string
	PostRepository model.PostRepository
}

// GetAllByUserProfile implements model.PostService
func (service *postService) GetAllByUserProfile(ctx context.Context, userProfileID uuid.UUID) ([]model.Post, error) {
	posts, err := service.PostRepository.GetAllByUserProfile(ctx, userProfileID)
	return posts, err
}

// PostServiceConfig is the parameter object for creating Post Services
type PostServiceConfig struct {
	BucketURL      string
	PostRepository model.PostRepository
}

// UploadPost implements model.PostService
func (service *postService) UploadPost(ctx context.Context, userProfileID uuid.UUID, caption string, image string) error {
	imageLink := service.bucketURL + image

	post := model.Post{
		UserProfileID: userProfileID,
		Caption:       caption,
		Image:         imageLink,
	}

	err := service.PostRepository.Create(ctx, post)
	return err
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
func NewPostService(config PostServiceConfig) model.PostService {
	return &postService{
		bucketURL:      config.BucketURL,
		PostRepository: config.PostRepository,
	}
}
