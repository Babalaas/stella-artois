package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostComment outlines the database entity for a comment on a post
type PostComment struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserProfileID uuid.UUID `json:"user_profile_id" gorm:"type:uuid;not null"`
	PostID        uuid.UUID `json:"post_id" gorm:"type:uuid;not null"`
	DateCreated   time.Time `json:"date_created" gorm:"type:timestamp with time zone;not null"`
	Content       string    `json:"content" gorm:"type:timestamp with time zone;not null"`
}

type PostCommentRepository interface {
	Create(ctx context.Context, comment *PostComment) (PostComment, error)
	Delete(ctx context.Context, comment *PostComment) error
}

type PostCommentService interface {
	Create(ctx context.Context, comment *PostComment) (PostComment, error)
	Delete(ctx context.Context, comment *PostComment) error
}

// BeforeCreate is the hook called before a PostComment is inserted into the database
// Used to initialize default values
func (comment *PostComment) BeforeCreate(db *gorm.DB) error {
	comment.ID = uuid.New()
	comment.DateCreated = time.Now().Local()
	return nil
}
