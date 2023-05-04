package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Friendship defines the friendship entity in the db
type Friendship struct {
	RequestUserProfileID  uuid.UUID `json:"request_user_profile_id" gorm:"type:uuid;not null"`
	ResponseUserProfileID uuid.UUID `json:"response_user_profile_id" gorm:"type:uuid;not null"`
	Status                string    `json:"status" gorm:"type:varchar(10);not null"`
	DateUpdated           time.Time `json:"date_updated" gorm:"type:timestamp with time zone;not null"`
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
	GetFriendsPosts(ctx context.Context, userProfileID uuid.UUID) ([]Post, error)
	RequestFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	AcceptFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	RemoveFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	FindFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) (Friendship, error)
	GetPendingFriendships(ctx context.Context, userProfileID uuid.UUID) ([]UserProfile, error)
	SearchNonFriends(ctx context.Context, userProfileID uuid.UUID, query string) ([]UserProfile, error)
}

// FriendshipService defines the usecases involving friendships
type FriendshipService interface {
	GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]UserProfile, error)
	RequestFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	AcceptFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	RemoveFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	GetFriendRequests(ctx context.Context, userProfileID uuid.UUID) ([]UserProfile, error)
	SearchNonFriends(ctx context.Context, userProfileID uuid.UUID, query string) ([]UserProfile, error)
}
