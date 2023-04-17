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
	UserProfileID uuid.UUID `json:"user_profile_id" binding:"required"`
	PostID        uuid.UUID `json:"post_id" binding:"required"`
	Content       string    `json:"content" binding:"required"`
}

type deletePostCommentRequest struct {
	ID            uuid.UUID `json:"id" binding:"required"`
	UserProfileID uuid.UUID `json:"user_profile_id" binding:"required"`
	PostID        uuid.UUID `json:"post_id" binding:"required"`
	DateCreated   time.Time `json:"date_created" binding:"required"`
	Content       string    `json:"content" binding:"required"`
}

// CreatePostComment is the HTTP handler to create one new PostComment
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

	resComment, createErr := handler.CommentService.Create(c.Request.Context(), reqComment)

	if createErr != nil {
		log.Panicf("Failed to create post comment: %v\n", createErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": createErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": resComment,
	})
}

// DeletePostComment is HTTP handler to delete one post comment by id
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

	err := handler.CommentService.Delete(c.Request.Context(), reqComment)

	if err != nil {
		log.Panicf("Unable to delete post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": "comment deleted successfully",
	})
}

// GetAllComments takes a post id and returns all post comments associated with that post
func (handler *Handler) GetAllComments(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))

	resComments, resErr := handler.CommentService.GetAll(c.Request.Context(), uid)

	if resErr != nil {
		log.Panicf("Unable to get comemnts")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get comments for post"})
		return
	}

	var trimmedComments []struct {
		UserProfileID string `json:"user_profile_id"`
		DisplayName   string `json:"display_name"`
		Content       string `json:"content"`
	}
	for _, comment := range resComments {
		displayName, err := handler.UserProfileService.GetDisplayName(c.Request.Context(), comment.UserProfileID)

		if err != nil {
			log.Fatal("Could not find user profile by id")
		}
		trimmedComment := struct {
			UserProfileID string `json:"user_profile_id"`
			DisplayName   string `json:"display_name"`
			Content       string `json:"content"`
		}{
			UserProfileID: comment.UserProfileID.String(),
			DisplayName:   displayName,
			Content:       comment.Content,
		}
		trimmedComments = append(trimmedComments, trimmedComment)
	}

	// Return list of users with only the desired fields
	c.JSON(http.StatusOK, trimmedComments)
}
