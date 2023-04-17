package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *Handler) GenerateFeed(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()

	feedPosts, err := handler.FeedService.GenerateFeed(uid, ctx)

	if err != nil {
		log.Panicf("Unable to get feed")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cant create feeds!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feed": feedPosts,
	})
}
