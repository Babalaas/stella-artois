package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserProfileID uuid.UUID `json:"user_profile_id" gorm:"type:uuid;not null"`
	Day           time.Time `json:"day" gorm:"type:date;not null"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
}

type CollectionRepoistory interface {
	Create(ctx context.Context, collection Collection) error
}

type CollectionService interface {
	CreateEmptyCollection(ctx context.Context, collection Collection) error
}
