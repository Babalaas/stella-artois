package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Post entity
type Post struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserProfileID uuid.UUID `gorm:"type:uuid;not null"`
	CollectionID  uuid.UUID `gorm:"type:uuid"`
	Caption       string    `gorm:"type:varchar(255);not null"`
	DateCreated   time.Time `gorm:"type:timestamp without time zone;not null"`
	Image         string    `gorm:"type:text;not null"`
	Image2        string    `gorm:"type:varchar(255)"`
	ReactionCount int       `gorm:"type:int4;not null"`
	InCollection  bool      `gorm:"type:boolean; not null"`
}

// PostService interface definition
type PostService interface {
	GetByID(ctx context.Context, uid uuid.UUID) (post Post, err error)
	AddToCollection(ctx context.Context, post *Post) (err error)
	UploadPost(ctx context.Context, userProfileID uuid.UUID, caption string, image string) error
	GetAllByUserProfile(ctx context.Context, userProfileID uuid.UUID) ([]Post, error)
}

// PostRepository interface definition
type PostRepository interface {
	GetByID(ctx context.Context, uid uuid.UUID) (post Post, err error)
	Create(ctx context.Context, post Post) error
	GetAllByUserProfile(ctx context.Context, userProfileID uuid.UUID) ([]Post, error)
}

// BeforeCreate is a hook called to initialize post fields to default values
func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.ID = uuid.New()
	post.CollectionID = uuid.Nil
	post.DateCreated = time.Now().Local()
	post.ReactionCount = 0
	post.InCollection = false
	return nil
}
