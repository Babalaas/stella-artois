package handler

import (
	"babalaas/stella-artois/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /{id}
func (handler *Handler) GetPostById(c *gin.Context) {
	post, exists := c.Get("post")
	if !exists {
		log.Panicf("Unable to extract post from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post record not found!"})
		return
	}
	uid := post.(*models.Post).ID
	ctx := c.Request.Context()
	post, err := handler.PostService.Get(ctx, uid)

	if err != nil {
		log.Panicf("Unable to find post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
