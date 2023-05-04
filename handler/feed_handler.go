package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenerateFeed returns posts to populate a users feed by id
func (handler *Handler) GenerateFeed(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()

	feedPosts, err := handler.FeedService.GenerateFeed(ctx, uid)

	if err != nil {
		log.Panicf("Unable to get feed")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cant create feeds!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feed": feedPosts,
	})
}

// GetPostsInCollection returns posts in a collection with a feed structure
func (handler *Handler) GetPostsInCollection(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()

	feedPosts, err := handler.FeedService.GenerateCollectionFeed(ctx, uid)

	if err != nil {
		log.Panicf("Unable to get feed")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cant create feeds!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feed": feedPosts,
	})
}

// GetPostsInCollection returns posts in a collection with a feed structure
func (handler *Handler) GetPostsByUserID(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()

	feedPosts, err := handler.FeedService.GenerateUserPostsFeed(ctx, uid)

	if err != nil {
		log.Panicf("Unable to get feed")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cant create feeds!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feed": feedPosts,
	})
}
