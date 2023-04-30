package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type collectionRepository struct {
	DB *gorm.DB
}

// DeleteByID implements model.CollectionRepository
func (repo *collectionRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	var collection model.Collection

	repo.DB.Where("id = ?", id).First(&collection)

	if collection.ID == uuid.Nil {
		log.Panic("Could not find collection with given ID")
	}
	result := repo.DB.Where("id = ?", id).Delete(id)

	if result.Error != nil {
		log.Panic("Could not delete collection.")
		return result.Error
	}

	return nil
}

// Create implements model.CollectionRepoistory
func (repo *collectionRepository) Create(ctx context.Context, collection model.Collection) error {
	result := repo.DB.Create(&collection).Error
	return result
}

// NewCollectionRepository is the factory function for created collection repos
func NewCollectionRepository(db *gorm.DB) model.CollectionRepository {
	return &collectionRepository{
		DB: db,
	}
}
