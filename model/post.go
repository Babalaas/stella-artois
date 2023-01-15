package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Post entity
type Post struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProfileID    uuid.UUID `gorm:"type:uuid;not null"`
	CollectionID uuid.UUID `gorm:"type:uuid"`
	Caption      string    `gorm:"type:varchar(255);not null"`
	Location     string    `gorm:"type:varchar(255);not null"`
	Created      time.Time `gorm:"type:timestamp without time zone;not null"`
	Image        string    `gorm:"type:varchar(255);not null"`
	Image2       string    `gorm:"type:varchar(255)"`
	DrinkNumber  int       `gorm:"type:int2"`
	LikeCount    int       `gorm:"type:int4;not null"`
}

// PostService interface definition
type PostService interface {
	GetByID(ctx context.Context, uid uuid.UUID) (post Post, err error)
	AddToCollection(ctx context.Context, post *Post) (err error)
}

// PostRepository interface definition
type PostRepository interface {
	GetAll(ctx context.Context) (posts []Post, err error)
	GetByID(ctx context.Context, uid uuid.UUID) (post Post, err error)
	Create(ctx context.Context, post *Post) (err error)
	Update(ctx context.Context, post *Post) (err error)
	Delete(ctx context.Context, uid uuid.UUID) (err error)
}

// BeforeCreate is a hook called to initialize post fields to default values
func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.ID = uuid.New()
	post.CollectionID = uuid.Nil
	post.Created = time.Now().Local()
	post.Created.Format(time.RFC3339)
	post.LikeCount = 0
	return nil
}
