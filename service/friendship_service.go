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

// SearchNonFriends implements model.FriendshipService
func (service *friendshipService) SearchNonFriends(ctx context.Context, userProfileID uuid.UUID, query string) ([]model.Friend, error) {
	friends, err := service.friendshipRepo.SearchNonFriends(ctx, userProfileID, query)

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

// GetFriendRequests implements model.FriendshipService
func (service *friendshipService) GetFriendRequests(ctx context.Context, userProfileID uuid.UUID) ([]model.UserProfile, error) {
	friendships, err := service.friendshipRepo.GetPendingFriendships(ctx, userProfileID)
	if err != nil {
		log.Panic("Could not find friend requests")
	}

	return friendships, err
}

// RemoveFriend implements model.FriendshipService
func (service *friendshipService) RemoveFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	err := service.friendshipRepo.RemoveFriendship(ctx, userProfileID, friendID)

	if err != nil {
		log.Panic("Could not remove friend")
	}

	return err
}

// RequestFriend implements model.FriendshipService
func (service *friendshipService) RequestFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	err := service.friendshipRepo.RequestFriendship(ctx, userProfileID, friendID)

	if err != nil {
		log.Panic("Could not generate friend request")
	}

	return err
}

// RespondToFriendshipRequest implements model.FriendshipService
func (service *friendshipService) AcceptFriend(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	_, err := service.friendshipRepo.FindFriendship(ctx, userProfileID, friendID)
	if err != nil {
		log.Panic("Could not find friendship")
	}

	err = service.friendshipRepo.AcceptFriendship(ctx, userProfileID, friendID)
	if err != nil {
		log.Panic("Could not accept friendship")
	}
	return err
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
