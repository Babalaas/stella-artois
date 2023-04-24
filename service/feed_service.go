package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

// FeedPost represents one post in a user's feed
type FeedPost struct {
	Author      author       `json:"author"`
	Post        post         `json:"post"`
	TopComments []topComment `json:"top_comments"`
}

type author struct {
	UserProfileID uuid.UUID `json:"request_user_profile_id"`
	DisplayName   string    `json:"display_name"`
	ProfilePic    string    `json:"profile_pic"`
}

type post struct {
	ID            uuid.UUID `json:"id"`
	Image         string    `json:"image"`
	Image2        string    `json:"image_2"`
	Caption       string    `json:"caption"`
	ReactionCount int       `json:"reaction_count"`
}

type topComment struct {
	UserProfileID uuid.UUID `json:"user_profile_id"`
	DisplayName   string    `json:"display_name"`
	Content       string    `json:"content"`
}

// FeedService generates the feed for a user
type FeedService interface {
	GenerateFeed(ctx context.Context, userProfileID uuid.UUID) ([]FeedPost, error)
}

type feedService struct {
	userProfileRepository model.UserProfileRepository
	postRepository        model.PostRepository
	commentRepository     model.CommentRepository
	reactionRepository    model.ReactionRepository
	friendshipRepository  model.FriendshipRepository
}

// FeedServiceConfig acts a paramter object for creating new FeedServices
type FeedServiceConfig struct {
	UserProfileRepository model.UserProfileRepository
	PostRepository        model.PostRepository
	CommentRepository     model.CommentRepository
	ReactionRepository    model.ReactionRepository
	FriendshipRepository  model.FriendshipRepository
}

// GenerateFeed implements FeedService
func (service *feedService) GenerateFeed(ctx context.Context, userProfileID uuid.UUID) ([]FeedPost, error) {
	var feedPosts []FeedPost

	// GetAllFriendsPosts
	posts, err := service.friendshipRepository.GetFriendsPosts(ctx, userProfileID)

	if err != nil {
		log.Panic("Could not get firends posts in feed service")
		return nil, err
	}

	// Build FeedPost structs
	for i := range posts {
		iPost, err := service.postRepository.GetByID(ctx, posts[i].ID)

		if err != nil {
			log.Fatal("OOF")
		}

		userProfile, err := service.userProfileRepository.FindByID(ctx, iPost.UserProfileID)

		if err != nil {
			log.Fatal("OOF")
		}

		comments, err := service.commentRepository.GetRecent(ctx, posts[i].ID, 2)

		if err != nil {
			log.Fatal("OOF")
		}

		var topComments []topComment
		for _, comment := range comments {
			commentAuthor, _ := service.userProfileRepository.FindByID(ctx, comment.UserProfileID)
			newTopComment := topComment{
				UserProfileID: comment.UserProfileID,
				DisplayName:   commentAuthor.DisplayName,
				Content:       comment.Content,
			}
			topComments = append(topComments, newTopComment)
		}

		feedPost := FeedPost{
			Author: author{
				UserProfileID: userProfile.ID,
				DisplayName:   userProfile.DisplayName,
				ProfilePic:    userProfile.ProfilePic,
			},
			Post: post{
				ID:            iPost.ID,
				Image:         iPost.Image,
				Image2:        iPost.Image2,
				Caption:       iPost.Caption,
				ReactionCount: iPost.ReactionCount,
			},
			TopComments: topComments,
		}

		feedPosts = append(feedPosts, feedPost)
	}

	return feedPosts, err
}

// NewFeedService is the facory function for creating a FeedService
func NewFeedService(config FeedServiceConfig) FeedService {
	return &feedService{
		userProfileRepository: config.UserProfileRepository,
		postRepository:        config.PostRepository,
		commentRepository:     config.CommentRepository,
		reactionRepository:    config.ReactionRepository,
		friendshipRepository:  config.FriendshipRepository,
	}
}
