package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

type FeedPost struct {
	Author      author       `json:"author"`
	Post        post         `json:"post"`
	TopComments []topComment `json:"top_comments"`
}

type author struct {
	UserProfileID uuid.UUID `json:"user_profile_id"`
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
	ID            uuid.UUID `json:"id"`
	UserProfileID uuid.UUID `json:"user_profile_id"`
	DisplayName   string    `json:"display_name"`
	Content       string    `json:"content"`
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
				ID:            comment.ID,
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

func NewFeedService(config FeedServiceConfig) FeedService {
	return &feedService{
		userProfileRepository: config.UserProfileRepository,
		postRepository:        config.PostRepository,
		commentRepository:     config.CommentRepository,
		reactionRepository:    config.ReactionRepository,
		friendshipRepository:  config.FriendshipRepository,
	}
}
