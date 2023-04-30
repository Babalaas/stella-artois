package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Collection defines the struct to represent the b_collection entity
type Collection struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserProfileID uuid.UUID `json:"user_profile_id" gorm:"type:uuid;not null"`
	Day           time.Time `json:"day" gorm:"type:date;not null"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
}

// CollectionRepository defines the database queries need to interact with collections
type CollectionRepository interface {
	Create(ctx context.Context, collection Collection) error
}

// CollectionService defines the use cases related to collections
type CollectionService interface {
	CreateEmptyCollection(ctx context.Context, collection Collection) error
}

// BeforeCreate is a hook called to initialize collection fields to default values
func (collection *Collection) BeforeCreate(db *gorm.DB) error {
	collection.ID = uuid.New()
	return nil
}

// TableName defines the name of the Collection table
func (Collection) TableName() string {
	return "b_collection"
}
