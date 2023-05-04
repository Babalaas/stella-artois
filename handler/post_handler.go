package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type uploadPostRequest struct {
	UserProfileID uuid.UUID `json:"user_profile_id"`
	Caption       string    `json:"caption"`
	Image         string    `json:"image"`
}

// GetPostByID cleans the id parameter and calls the PostService to get a post by based on the parameter (GET /{id})
func (handler *Handler) GetPostByID(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	post, err := handler.PostService.GetByID(ctx, uid)

	if err != nil {
		log.Println("Unable to find post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

// UploadPost handles HTTP request to create a new post
func (handler *Handler) UploadPost(c *gin.Context) {
	var req uploadPostRequest

	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		log.Printf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.PostService.UploadPost(c, req.UserProfileID, req.Caption, req.Image)

	if err != nil {
		log.Println("Unable to upload post")
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// GetPostsByUser cleans the id parameter and calls the PostService to get a post by based on the parameter (GET /{id})
func (handler *Handler) GetPostsByUser(c *gin.Context) {
	reqID := c.Query("user_profile_id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	posts, err := handler.PostService.GetAllByUserProfile(ctx, uid)

	if err != nil {
		log.Println("Unable to find posts")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
