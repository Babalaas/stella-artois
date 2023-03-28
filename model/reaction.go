package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostReaction reperesents the post reaction entity
type PostReaction struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primarykey"`
	UserProfileID uuid.UUID `json:"user_profile_id" gorm:"type:uuid;not null"`
	PostID        uuid.UUID `json:"post_id" gorm:"type:uuid;not null"`
	DateCreated   time.Time `json:"date_created" gorm:"type:timestamp with time zone;not null"`
	ReactionID    int       `json:"reaction_id" gorm:"type:int4;not null"`
}

// ReactionRepository defines the reaction entity's database interactions
type ReactionRepository interface {
	Create(ctx context.Context, reaction *PostReaction) (PostReaction, error)
}

// ReactionService defines the use cases related to reactions
type ReactionService interface {
	ReactToPost(ctx context.Context, reaction *PostReaction) (PostReaction, error)
}

// BeforeCreate is the hook called before a PostReactuib is inserted into the database.
// Used to initialize default values
func (reaction *PostReaction) BeforeCreate(db *gorm.DB) error {
	reaction.ID = uuid.New()
	reaction.DateCreated = time.Now().Local()
	return nil
}
