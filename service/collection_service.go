package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"
)

type COLSConfig struct {
	collectionRepo model.CollectionRepoistory
}

type collectionService struct {
	collectionRepo model.CollectionRepoistory
}

// CreateEmptyCollection implements model.CollectionService
func (service *collectionService) CreateEmptyCollection(ctx context.Context, collection model.Collection) error {
	err := service.collectionRepo.Create(ctx, collection)
	if err != nil {
		log.Fatal("Can not create empty collection")
	}
	return err
}

func NewCollectionService(config COLSConfig) model.CollectionService {
	return &collectionService{
		collectionRepo: config.collectionRepo,
	}
}
