package service

import (
	"babalaas/stella-artois/model"
	"context"
)

type FeedPost struct {
	UserProfile model.UserProfile
	Post        model.Post
	Comments    []model.PostComment
}

type FeedService interface {
	GenerateFeed(userProfile model.UserProfile, ctx context.Context) ([]FeedPost, error)
}

type feedService struct {
	userProfileRepository model.UserProfileRepository
	postRepository        model.PostRepository
	commentRepository     model.CommentRepository
	reactionRepository    model.ReactionRepository
}

// GenerateFeed implements FeedService
func (*feedService) GenerateFeed(userProfile model.UserProfile, ctx context.Context) ([]FeedPost, error) {
	panic("unimplemented")
}

type FeedServiceConfig struct {
	UserProfileRepository model.UserProfileRepository
	PostRepository        model.PostRepository
	CommentRepository     model.CommentRepository
	ReactionRepository    model.ReactionRepository
}

func NewFeedService(config FeedServiceConfig) FeedService {
	return &feedService{
		userProfileRepository: config.UserProfileRepository,
		postRepository:        config.PostRepository,
		commentRepository:     config.CommentRepository,
		reactionRepository:    config.ReactionRepository,
	}
}
