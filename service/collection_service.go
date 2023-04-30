package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

// CollectionServiceConfig is the parameter object for creating a Collection Service
type CollectionServiceConfig struct {
	CollectionRepo model.CollectionRepository
}

type collectionService struct {
	collectionRepo model.CollectionRepository
}

// Delete implements model.CollectionService
func (service *collectionService) Delete(ctx context.Context, id uuid.UUID) error {
	err := service.collectionRepo.DeleteByID(ctx, id)
	return err
}

// CreateEmptyCollection implements model.CollectionService
func (service *collectionService) CreateEmptyCollection(ctx context.Context, collection model.Collection) error {
	err := service.collectionRepo.Create(ctx, collection)
	if err != nil {
		log.Fatal("Can not create empty collection")
	}
	return err
}

// NewCollectionService is the factory function for created collection services
func NewCollectionService(config CollectionServiceConfig) model.CollectionService {
	return &collectionService{
		collectionRepo: config.CollectionRepo,
	}
}
