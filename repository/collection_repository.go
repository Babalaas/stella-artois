package repository

import (
	"babalaas/stella-artois/model"
	"context"

	"gorm.io/gorm"
)

type collectionRepository struct {
	DB *gorm.DB
}

// Create implements model.CollectionRepoistory
func (repo *collectionRepository) Create(ctx context.Context, collection model.Collection) error {
	result := repo.DB.Create(&collection).Error
	return result
}

func NewCollectionRepository(db *gorm.DB) model.CollectionRepoistory {
	return &collectionRepository{
		DB: db,
	}
}
