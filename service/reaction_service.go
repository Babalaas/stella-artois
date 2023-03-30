package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"
)

type reactionService struct {
	ReactionRepo model.ReactionRepository
}

// ReactToPost implements model.ReactionService
func (service *reactionService) ReactToPost(ctx context.Context, reaction *model.PostReaction) (model.PostReaction, error) {
	_, err := service.ReactionRepo.Create(ctx, reaction)

	if err != nil {
		log.Panic("BAD")
	}
	return *reaction, err
}

// RSConfig is the parameter object for creating new ReactionService structs
type RSConfig struct {
	ReactionRepo model.ReactionRepository
}

// NewReactionService is the factory function for creating new ReactionService structs
func NewReactionService(config *RSConfig) model.ReactionService {
	return &reactionService{
		ReactionRepo: config.ReactionRepo,
	}
}
