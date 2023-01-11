package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Post struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Profile_ID    uuid.UUID `gorm:"type:uuid;not null"`
	Collection_ID uuid.UUID `gorm:"type:uuid"`
	Caption       string    `gorm:"type:varchar(255);not null"`
	Location      string    `gorm:"type:varchar(255);not null"`
	Created       time.Time `gorm:"type:timestamp without time zone;not null"`
	Image         string    `gorm:"type:varchar(255);not null"`
	Image2        string    `gorm:"type:varchar(255)"`
	Drink_Number  int       `gorm:"type:int2"`
	Like_Count    int       `gorm:"type:int4;not null"`
}

type PostService interface {
	Get(ctx context.Context, uid uuid.UUID) (post Post, err error)
	AddToCollection(ctx context.Context, post *Post) (err error)
}

type PostRepository interface {
	GetAll(ctx context.Context) (posts []Post, err error)
	GetById(ctx context.Context, uid uuid.UUID) (post Post, err error)
	Create(ctx context.Context, post *Post) (err error)
	Update(ctx context.Context, post *Post) (err error)
	Delete(ctx context.Context, uid uuid.UUID) (err error)
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.ID = uuid.New()
	post.Collection_ID = uuid.Nil
	post.Created = time.Now().Local()
	post.Created.Format(time.RFC3339)
	post.Like_Count = 0
	return nil
}
