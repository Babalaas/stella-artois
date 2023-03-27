package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
