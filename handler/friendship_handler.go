package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type friendshipRequest struct {
	RequesterID uuid.UUID
	ResponderID uuid.UUID
}

type friendshipResponseRequest struct {
	RequesterID uuid.UUID
	ResponderID uuid.UUID
	decision    int
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

func (handler *Handler) RequestFriend(c *gin.Context) {
	var request friendshipRequest

	if bindErr := c.ShouldBind(&request); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.RequestFriend(c, request.RequesterID, request.ResponderID)

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

func (handler *Handler) RepondToFriendRequest(c *gin.Context) {
	var req friendshipResponseRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.RespondToFriendshipRequest(c, req.RequesterID, req.ResponderID, req.decision)

	if err != nil {
		log.Panic("Failed to respond to request")
	}

	c.JSON(http.StatusOK, gin.H{
		"frienship": "updated",
	})
}

func (handler *Handler) RemoveFriend(c *gin.Context) {
	var req friendshipResponseRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind friendshio JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.FriendshipService.RemoveFriend(c, req.RequesterID, req.ResponderID)

	if err != nil {
		log.Panic("Failed to remove friend")
	}

	c.JSON(http.StatusOK, gin.H{
		"frienship": "removed",
	})
}
