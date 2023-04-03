package handler

import (
	"babalaas/stella-artois/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postReactionRequest struct {
	UserProfileID uuid.UUID `json:"user_profile_id" binding:"required"`
	PostID        uuid.UUID `json:"post_id" binding:"required"`
	ReactionID    int       `json:"reaction_id" binding:"required"`
}

// ReactToPost creates a new post reaction
func (handler *Handler) ReactToPost(c *gin.Context) {
	var req postReactionRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	reqReaction := &model.PostReaction{
		UserProfileID: req.UserProfileID,
		PostID:        req.PostID,
		ReactionID:    req.ReactionID,
	}

	resReaction, createErr := handler.ReactionService.ReactToPost(c.Request.Context(), reqReaction)

	if createErr != nil {
		log.Panicf("Failed to create post comment: %v\n", createErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": createErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": resReaction,
	})
}
