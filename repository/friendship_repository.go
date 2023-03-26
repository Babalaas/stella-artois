package repository

import (
	"babalaas/stella-artois/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type friendshipRepository struct {
	DB *gorm.DB
}

// GetAllFriends implements model.FriendshipRepository
func (repo *friendshipRepository) GetAllFriends(ctx context.Context, userProfile *model.UserProfile) ([]model.UserProfile, error) {
	var friends []model.UserProfile
	userProfileID := userProfile.ID
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
