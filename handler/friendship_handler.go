package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type friendshipRequest struct {
	RequestUserProfileID  uuid.UUID `json:"request_user_profile_id" binding:"required"`
	ResponseUserProfileID uuid.UUID `json:"response_user_profile_id" binding:"required"`
}

// GetAllFriends is the HTTP handler to return the passed userProfileID's friends in a list
func (handler *Handler) GetAllFriends(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	friends, err := handler.FriendshipService.GetAllFriends(ctx, uid)

	if err != nil {
		log.Panicf("Friendship Handler: Unable to find friends post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Friends not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}

// RequestFriend is the HTTP handler to request a friendship between two userProfileIDs
func (handler *Handler) RequestFriend(c *gin.Context) {
	var request friendshipRequest

	if bindErr := c.ShouldBind(&request); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.RequestFriend(c, request.RequestUserProfileID, request.ResponseUserProfileID)

	if err != nil {
		log.Panicf("Failed to register user profile: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"frienship": "requested",
	})
}

// AcceptFriend is the HTTP handler to accept a friendship between two userProfileIDs
func (handler *Handler) AcceptFriend(c *gin.Context) {
	var req friendshipRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.AcceptFriend(c, req.RequestUserProfileID, req.ResponseUserProfileID)

	if err != nil {
		log.Panic("Failed to respond to request")
	}

	c.JSON(http.StatusOK, gin.H{
		"frienship": "updated",
	})
}

// RemoveFriend is the HTTP handler to reject or delete a friendship between two userProfileIDs
func (handler *Handler) RemoveFriend(c *gin.Context) {
	var req friendshipRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.RemoveFriend(c, req.RequestUserProfileID, req.ResponseUserProfileID)

	if err != nil {
		log.Panic("Failed to remove friend")
	}

	c.JSON(http.StatusOK, gin.H{
		"frienship": "removed",
	})
}

// SearchNonFriends returns a list of users wiith displayanmes that contain the query text and are not the passed user or friends
func (handler *Handler) SearchNonFriends(c *gin.Context) {
	query := c.Query("q")
	reqID := c.Query("id")
	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	userProfiles, err := handler.FriendshipService.SearchNonFriends(ctx, uid, query)

	if err != nil {
		log.Printf("Failed to seach for user profiles: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": userProfiles,
	})
}
