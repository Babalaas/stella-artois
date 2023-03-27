package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserProfile entity
type UserProfile struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	DisplayName string    `json:"display_name" gorm:"type:varchar(30);not null"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(15);not null"`
	LastName    string    `json:"last_name" gorm:"type:varchar(30);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null"`
	Phone       string    `json:"phone" gorm:"type:varchar(15);not null"`
	Birthdate   time.Time `json:"birthdate" gorm:"type:date;not null"`
	Password    string    `json:"password" gorm:"type:text"`
	ProfilePic  string    `json:"profile_pic" gorm:"type:text;not null"`
}

// UserProfileService definition
type UserProfileService interface {
	Register(ctx context.Context, userProfile *UserProfile) (UserProfile, error)
	LogIn(ctx context.Context, userProfile *UserProfile) (UserProfile, error)
}

// UserProfileRepository definition
type UserProfileRepository interface {
	Create(ctx context.Context, userProfile *UserProfile) (UserProfile, error)
	FindByDisplayName(ctx context.Context, displayName string) (UserProfile, error)
}

// BeforeCreate is a hook called to initialize user_profile fields to default values
func (userProfile *UserProfile) BeforeCreate(db *gorm.DB) error {
	userProfile.ID = uuid.New()
	return nil
}
