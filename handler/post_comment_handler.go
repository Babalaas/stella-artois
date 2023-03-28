package handler

import (
	"babalaas/stella-artois/model"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPostCommentRequest struct {
	UserProfileID uuid.UUID
	PostID        uuid.UUID
	Content       string
}

type deletePostCommentRequest struct {
	ID            uuid.UUID
	UserProfileID uuid.UUID
	PostID        uuid.UUID
	DateCreated   time.Time
	Content       string
}

func (handler *Handler) CreatePostComment(c *gin.Context) {
	var req createPostCommentRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	reqComment := &model.PostComment{
		UserProfileID: req.UserProfileID,
		PostID:        req.PostID,
		Content:       req.Content,
	}

	resComment, createErr := handler.PostCommentService.Create(c.Request.Context(), reqComment)

	if createErr != nil {
		log.Panicf("Failed to register user profile: %v\n", createErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": createErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": resComment,
	})
}

func (handler *Handler) DeletePostComment(c *gin.Context) {
	var req deletePostCommentRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	reqComment := &model.PostComment{
		ID:            req.ID,
		UserProfileID: req.UserProfileID,
		PostID:        req.PostID,
		DateCreated:   req.DateCreated,
		Content:       req.Content,
	}

	err := handler.PostCommentService.Delete(c.Request.Context(), reqComment)

	if err != nil {
		log.Panicf("Unable to delete post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": "comment deleted successfully",
	})
}
