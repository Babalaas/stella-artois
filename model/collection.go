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

// CollectionPost defines the struc t representing the associative table between b_collection and post
type CollectionPost struct {
	CollectionID uuid.UUID `json:"collection_id" gorm:"type:uuid;column:b_collection_id;not null"`
	PostID       uuid.UUID `json:"post_id" gorm:"type:uuid;not null"`
	DateAdded    time.Time `gorm:"type:timestamp without time zone;not null"`
}

// CollectionRepository defines the database queries need to interact with collections
type CollectionRepository interface {
	CreateCollectionPost(ctx context.Context, postID uuid.UUID, collectionID uuid.UUID) error
	Create(ctx context.Context, collection Collection) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
	GetAllByUserProfileID(ctx context.Context, userProfileID uuid.UUID) ([]Collection, error)
	UpdateCollection(ctx context.Context, collection Collection) error
}

// CollectionService defines the use cases related to collections
type CollectionService interface {
	AddPostToCollection(ctx context.Context, postID uuid.UUID, collectionID uuid.UUID) error
	CreateEmptyCollection(ctx context.Context, collection Collection) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetUserCollections(ctx context.Context, userProfileID uuid.UUID) ([]Collection, error)
	UpdateCollection(ctx context.Context, collection Collection) error
}

// BeforeCreate is a hook called to initialize collection fields to default values
func (collection *Collection) BeforeCreate(db *gorm.DB) error {
	collection.ID = uuid.New()
	return nil
}

func (post *CollectionPost) BeforeCreate(db *gorm.DB) error {
	post.DateAdded = time.Now().Local()
	return nil
}

// TableName defines the name of the Collection table
func (Collection) TableName() string {
	return "b_collection"
}
