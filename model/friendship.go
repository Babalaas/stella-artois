package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Friendship struct {
	requesterID  uuid.UUID `gorm:"type:uuid;not null"`
	responderID  uuid.UUID `gorm:"type:uuid;not null"`
	status       string    `gorm:"type:varchar(10);not null"`
	date_updated time.Time `gorm:"type:timestamp with time zone;not null"`
}

// BeforeCreate is a hook called to initialize Friendship fields to default values
func (friendship *Friendship) BeforeCreate(db *gorm.DB) error {
	friendship.status = "Requested"
	friendship.date_updated = time.Now().Local()
	return nil
}

type FriendshipRepository interface {
	GetAllFriends(ctx context.Context, userProfile *UserProfile) ([]UserProfile, error)
}
