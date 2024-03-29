package service

import (
	"babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
)

// CollectionServiceConfig is the parameter object for creating a Collection Service
type CollectionServiceConfig struct {
	CollectionRepo model.CollectionRepository
}

type collectionService struct {
	collectionRepo model.CollectionRepository
}

// AddPostToCollection implements model.CollectionService
func (service *collectionService) AddPostToCollection(ctx context.Context, postID uuid.UUID, collectionID uuid.UUID) error {
	err := service.collectionRepo.CreateCollectionPost(ctx, postID, collectionID)
	return err
}

// UpdateCollection implements model.CollectionService
func (service *collectionService) UpdateCollection(ctx context.Context, collection model.Collection) error {
	err := service.collectionRepo.UpdateCollection(ctx, collection)
	return err
}

// GetUserCollections implements model.CollectionService
func (service *collectionService) GetUserCollections(ctx context.Context, userProfileID uuid.UUID) ([]model.Collection, error) {
	collections, err := service.collectionRepo.GetAllByUserProfileID(ctx, userProfileID)
	return collections, err
}

// Delete implements model.CollectionService
func (service *collectionService) Delete(ctx context.Context, id uuid.UUID) error {
	err := service.collectionRepo.DeleteByID(ctx, id)
	return err
}

// CreateEmptyCollection implements model.CollectionService
func (service *collectionService) CreateEmptyCollection(ctx context.Context, collection model.Collection) error {
	err := service.collectionRepo.Create(ctx, collection)
	return err
}

// NewCollectionService is the factory function for created collection services
func NewCollectionService(config CollectionServiceConfig) model.CollectionService {
	return &collectionService{
		collectionRepo: config.CollectionRepo,
	}
}
