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

// SearchNonFriends implements model.FriendshipRepository
func (*friendshipRepository) SearchNonFriends(ctx context.Context, userProfileID uuid.UUID, query string) ([]model.UserProfile, error) {
	panic("unimplemented")
}

// GetPendingFriendships implements model.FriendshipRepository
func (repo *friendshipRepository) GetPendingFriendships(ctx context.Context, userProfileID uuid.UUID) ([]model.UserProfile, error) {
	var userProfiles []model.UserProfile
	err := repo.DB.Table("user_profile").
		Joins("INNER JOIN friendship ON friendship.request_user_profile_id = user_profile.id").
		Where("friendship.response_user_profile_id = ? AND friendship.status = ?", userProfileID, "requested").
		Find(&userProfiles).Error
	if err != nil {
		log.Panic("Could not get pending friendships")
		return userProfiles, err
	}
	return userProfiles, err
}

// FindFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) FindFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) (model.Friendship, error) {
	var friendship model.Friendship

	result := repo.DB.Where("(request_user_profile_id = ? AND response_user_profile_id = ?) OR (request_user_profile_id = ? AND response_user_profile_id = ?) AND status = ?", userProfileID, friendID, friendID, userProfileID, "accepted").First(&friendship)
	if result.Error != nil {
		return friendship, result.Error
	}
	return friendship, nil

}

// AcceptFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) AcceptFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	friendship, err := repo.FindFriendship(ctx, userProfileID, friendID)

	if err != nil {
		log.Panic("BAH HUMBUG", err)
		return err
	}

	friendship.Status = "accepted"
	friendship.DateUpdated = time.Now()

	// Save the updated Friendship record to the database
	err = repo.DB.Where("(request_user_profile_id = ? AND response_user_profile_id = ?) OR (request_user_profile_id = ? AND response_user_profile_id = ?) AND status = ?", userProfileID, friendID, friendID, userProfileID, "accepted").Save(&friendship).Error

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
	//err = repo.DB.Delete(friendship).Error
	err = repo.DB.Where("(request_user_profile_id = ? AND response_user_profile_id = ?) OR (request_user_profile_id = ? AND response_user_profile_id = ?) AND status = ?", userProfileID, friendID, friendID, userProfileID, "accepted").Delete(&friendship).Error

	if err != nil {
		log.Panic("could not delete friendship")
		return err
	}

	return err
}

// RequestFriendship implements model.FriendshipRepository
func (repo *friendshipRepository) RequestFriendship(ctx context.Context, userProfileID uuid.UUID, friendID uuid.UUID) error {
	friendship := model.Friendship{
		RequestUserProfileID:  userProfileID,
		ResponseUserProfileID: friendID,
		Status:                "requested",
		DateUpdated:           time.Now(),
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

	err := repo.DB.Table("post").
		Select("post.*").
		Joins("JOIN friendship ON (post.user_profile_id = friendship.response_user_profile_id OR post.user_profile_id = friendship.request_user_profile_id)").
		Where("(friendship.request_user_profile_id = ? OR friendship.response_user_profile_id = ?) AND friendship.status = ?", userProfileID, userProfileID, "accepted").
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
