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

// CreateCollectionPost implements model.CollectionRepository
func (repo *collectionRepository) CreateCollectionPost(ctx context.Context, postID uuid.UUID, collectionID uuid.UUID) error {
	collectionPost := model.CollectionPost{
		CollectionID: collectionID,
		PostID:       postID,
	}
	err := repo.DB.Create(&collectionPost).Error
	return err
}

// UpdateCollection implements model.CollectionRepository
func (repo *collectionRepository) UpdateCollection(ctx context.Context, collection model.Collection) error {
	err := repo.DB.Model(&collection).Where("id = ?", collection.ID).Updates(map[string]interface{}{
		"name": collection.Name,
		"day":  collection.Day,
	}).Error

	return err
}

// GetAllByUserProfileID implements model.CollectionRepository
func (repo *collectionRepository) GetAllByUserProfileID(ctx context.Context, userProfileID uuid.UUID) ([]model.Collection, error) {
	var collections []model.Collection

	err := repo.DB.Where("user_profile_id = ?", userProfileID).Find(&collections).Error

	if err != nil {
		log.Println("Error finding collections belong to user")
	}

	return collections, err
}

// DeleteByID implements model.CollectionRepository
func (repo *collectionRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	var collection model.Collection

	repo.DB.Where("id = ?", id).First(&collection)

	if collection.ID == uuid.Nil {
		log.Println("Could not find collection with given ID")
	}
	result := repo.DB.Where("id = ?", id).Delete(&collection)

	if result.Error != nil {
		log.Println("Could not delete collection.")
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
