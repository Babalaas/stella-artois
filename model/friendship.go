package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Friendship defines the friendship entity in the db
type Friendship struct {
	// requesterID  uuid.UUID `gorm:"type:uuid;not null"`
	// responderID  uuid.UUID `gorm:"type:uuid;not null"`
	Status      string    `gorm:"type:varchar(10);not null"`
	DateUpdated time.Time `gorm:"type:timestamp with time zone;not null"`
}

// BeforeCreate is a hook called to initialize Friendship fields to default values
func (friendship *Friendship) BeforeCreate(db *gorm.DB) error {
	friendship.Status = "Requested"
	friendship.DateUpdated = time.Now().Local()
	return nil
}

// FriendshipRepository defines how the application interacts with the db
type FriendshipRepository interface {
	GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]UserProfile, error)
}

// FriendshipService defines the usecases involving friendships
type FriendshipService interface {
	GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]UserProfile, error)
}
