package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type reactionRepository struct {
	DB *gorm.DB
}

// Create implements model.ReactionRepository
func (repo *reactionRepository) Create(ctx context.Context, reaction *model.PostReaction) (model.PostReaction, error) {
	result := repo.DB.Create(&reaction)

	if result.Error != nil {
		log.Panic("Could not create new User Profile.")
		return *reaction, result.Error
	}

	return *reaction, nil
}

func NewReactionRepository(db *gorm.DB) model.ReactionRepository {
	return &reactionRepository{
		DB: db,
	}
}
