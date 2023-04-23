package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type friendshipRepository struct {
	DB *gorm.DB
}

// FindFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) FindFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) (model.Friendship, error) {
	friendship := model.Friendship{}
	err := repo.DB.Find(friendship, "(requester_id = ? AND responder_id = ?) OR (requester_id = ? AND responder_id = ?)", userProfileID, friendID).Error
	if err != nil {
		return friendship, err
	}

	// Return the Friendship record if one was found
	return friendship, nil
}

// AcceptFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) AcceptFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	friendship, err := repo.FindFriendship(ctx, userProfileID, friendID)

	if err != nil {
		log.Panic("could not find friendship")
		return err
	}

	friendship.Status = "accepted"
	friendship.DateUpdated = time.Now()

	// Save the updated Friendship record to the database
	err = repo.DB.Save(friendship).Error

	if err != nil {
		log.Panic("could not update friendship")
		return err
	}

	return err
}

// RemoveFriendship can be used to reject a friendship request or remove a currenty friendship
func (repo *friendshipRepository) RemoveFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	friendship, err := repo.FindFriendship(ctx, userProfileID, friendID)

	if err != nil {
		log.Panic("could not find friendship")
		return err
	}

	friendship.Status = "rejected"
	friendship.DateUpdated = time.Now()

	// Save the updated Friendship record to the database
	err = repo.DB.Delete(friendship).Error

	if err != nil {
		log.Panic("could not delete friendship")
		return err
	}

	return err
}

// RequestFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) RequestFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	friendship := model.Friendship{
		RequesterID: userProfileID,
		ResponderID: friendID,
		Status:      "requested",
		DateUpdated: time.Now(),
	}

	result := repo.DB.Create(&friendship)

	if result.Error != nil {
		log.Panic("Could not create new User Profile.")
		return result.Error
	}

	return nil
}

// GetFriendsPosts implements model.FriendshipRepository
func (repo *friendshipRepository) GetFriendsPosts(ctx context.Context, userProfileID uuid.UUID) ([]model.Post, error) {
	var posts []model.Post

	err := repo.DB.Joins("JOIN user_profile ON post.user_profile_id = user_profile.id").
		Joins("JOIN friendship f1 ON post.user_profile_id = f1.response_user_profile_id").
		Joins("JOIN friendship f2 ON post.user_profile_id = f2.request_user_profile_id").
		Where("(f1.request_user_profile_id = ? AND f1.status = ?) OR (f2.response_user_profile_id = ? AND f2.status = ?)", userProfileID, "accepted", userProfileID, "accepted").
		Find(&posts).Error

	if err != nil {
		log.Panic("Could not get friends posts")
		return posts, err
	}

	return posts, nil
}

// GetAllFriends implements model.FriendshipRepository
func (repo *friendshipRepository) GetAllFriends(ctx context.Context, userProfileID uuid.UUID) ([]model.UserProfile, error) {
	var friends []model.UserProfile
	err := repo.DB.Table("user_profile").
		Select("user_profile.id, user_profile.display_name, user_profile.first_name, user_profile.last_name, user_profile.email, user_profile.phone, user_profile.birthdate, user_profile.profile_pic, friendship.status, friendship.date_updated").
		Joins("INNER JOIN friendship ON user_profile.id = friendship.request_user_profile_id OR user_profile.id = friendship.response_user_profile_id").
		Where("(friendship.request_user_profile_id = ? OR friendship.response_user_profile_id = ?) AND friendship.status = ?", userProfileID, userProfileID, "accepted").
		Find(&friends).Error

	if err != nil {
		log.Panic("Could not create a ust of friends")
		return nil, err
	}

	return friends, nil
}

// NewFriendshipRepository creates a new FriendshipRepository with the server's database instance
func NewFriendshipRepository(db *gorm.DB) model.FriendshipRepository {
	return &friendshipRepository{
		DB: db,
	}
}
