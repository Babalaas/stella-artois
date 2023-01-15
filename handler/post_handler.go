package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetPostByID cleans the id parameter and calls the PostService to get a post by based on the parameter (GET /{id})
func (handler *Handler) GetPostById(c *gin.Context) {
	reqId := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqId))
	ctx := c.Request.Context()
	post, err := handler.PostService.GetById(ctx, uid)

	if err != nil {
		log.Panicf("Unable to find post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
