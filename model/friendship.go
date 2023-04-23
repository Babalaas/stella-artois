package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Friendship defines the friendship entity in the db
type Friendship struct {
	RequesterID uuid.UUID `gorm:"type:uuid;not null"`
	ResponderID uuid.UUID `gorm:"type:uuid;not null"`
	Status      string    `gorm:"type:varchar(10);not null"`
	DateUpdated time.Time `gorm:"type:timestamp with time zone;not null"`
}

// Friend is used to communicate necessary fields about a friend to a user_profile
type Friend struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	DisplayName string    `json:"display_name" gorm:"type:varchar(30);not null"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(15);not null"`
	LastName    string    `json:"last_name" gorm:"type:varchar(30);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null"`
	Phone       string    `json:"phone" gorm:"type:varchar(15);not null"`
	ProfilePic  string    `json:"profile_pic" gorm:"type:text;not null"`
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
}

// FriendshipService defines the usecases involving friendships
type FriendshipService interface {
	GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]Friend, error)
	RequestFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
	RespondToFriendshipRequest(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID, decision int) error
	RemoveFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error
}
