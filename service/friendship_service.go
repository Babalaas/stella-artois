package service

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"github.com/google/uuid"
)

type friendshipService struct {
	friendshipRepo model.FriendshipRepository
}

// FSConfig is the paramter object used to create new Friendship Services
// Holds required dependencies
type FSConfig struct {
	FriendshipRepository model.FriendshipRepository
}

// GetAllFriends implements model.FriendshipService
func (service *friendshipService) GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]model.Friend, error) {
	friends, err := service.friendshipRepo.GetAllFriends(ctx, userProfileID)

	if err != nil {
		log.Fatal("Friendship Service: error getting all friends from repo")
		return nil, err
	}

	var trimmedFriends []model.Friend

	for _, friend := range friends {
		trimmedFriend := model.Friend{
			ID:          friend.ID,
			DisplayName: friend.DisplayName,
			FirstName:   friend.FirstName,
			LastName:    friend.LastName,
			Email:       friend.Email,
			Phone:       friend.Phone,
			ProfilePic:  friend.ProfilePic}
		trimmedFriends = append(trimmedFriends, trimmedFriend)
	}

	return trimmedFriends, err
}

// NewFriendshipService creates a Post Service with a PostRepository attribute
func NewFriendshipService(config FSConfig) model.FriendshipService {
	return &friendshipService{
		friendshipRepo: config.FriendshipRepository,
	}
}
