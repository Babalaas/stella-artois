package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	DisplayName string    `gorm:"type:varchar(30);not null"`
	FirstName   string    `gorm:"type:varchar(15);not null"`
	LastName    string    `gorm:"type:varchar(30);not null"`
	Email       string    `gorm:"type:varchar(255);not null"`
	Phone       string    `gorm:"type:varchar(15);not null"`
	Gender      string    `gorm:"type:varchar(25);not null"`
	Birthdate   time.Time `gorm:"type:date;not null"`
	Password    string    `gorm:"type:text"`
}

type UserProfileService interface {
	Register(ctx context.Context, userProfile *UserProfile) (err error)
}

type UserProfileRepository interface {
	Create(ctx context.Context, userProfile *UserProfile) (err error)
}

// BeforeCreate is a hook called to initialize user_profile fields to default values
func (userProfile *UserProfile) BeforeCreate(db *gorm.DB) error {
	userProfile.ID = uuid.New()
	return nil
}
