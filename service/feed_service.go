package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

type FeedPost struct {
	UserProfile model.UserProfile
	Post        model.Post
	Comments    []model.PostComment
}

type FeedService interface {
	GenerateFeed(userProfileID uuid.UUID, ctx context.Context) ([]FeedPost, error)
}

type feedService struct {
	userProfileRepository model.UserProfileRepository
	postRepository        model.PostRepository
	commentRepository     model.CommentRepository
	reactionRepository    model.ReactionRepository
	friendshipRepository  model.FriendshipRepository
}

type FeedServiceConfig struct {
	UserProfileRepository model.UserProfileRepository
	PostRepository        model.PostRepository
	CommentRepository     model.CommentRepository
	ReactionRepository    model.ReactionRepository
	FriendshipRepository  model.FriendshipRepository
}

// GenerateFeed implements FeedService
func (service *feedService) GenerateFeed(userProfileID uuid.UUID, ctx context.Context) ([]FeedPost, error) {
	var feedPosts []FeedPost

	// GetAllFriendsPosts
	posts, err := service.friendshipRepository.GetFriendsPosts(ctx, userProfileID)

	if err != nil {
		log.Panic("Could not get firends posts in feed service")
		return nil, err
	}

	// Build FeedPost structs
	for i := range posts {
		post, err := service.postRepository.GetByID(ctx, posts[i].ID)

		if err != nil {
			log.Fatal("OOF")
		}

		userProfile, err := service.userProfileRepository.FindByID(ctx, post.UserProfileID)

		if err != nil {
			log.Fatal("OOF")
		}

		comments, err := service.commentRepository.GetRecent(ctx, posts[i].ID, 2)

		if err != nil {
			log.Fatal("OOF")
		}

		feedPost := FeedPost{
			UserProfile: userProfile,
			Post:        post,
			Comments:    comments,
		}

		feedPosts = append(feedPosts, feedPost)
	}

	return feedPosts, err
}

func NewFeedService(config FeedServiceConfig) FeedService {
	return &feedService{
		userProfileRepository: config.UserProfileRepository,
		postRepository:        config.PostRepository,
		commentRepository:     config.CommentRepository,
		reactionRepository:    config.ReactionRepository,
		friendshipRepository:  config.FriendshipRepository,
	}
}
