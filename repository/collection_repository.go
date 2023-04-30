package repository

import (
	"babalaas/stella-artois/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type collectionRepository struct {
	DB *gorm.DB
}

// DeleteByID implements model.CollectionRepository
func (*collectionRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
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
