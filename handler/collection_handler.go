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

type createEmptyRequest struct {
	UserProfileID uuid.UUID `json:"user_profile_id" binding:"required"`
	Day           time.Time `json:"day" binding:"required"`
	Name          string    `json:"name" binding:"required"`
}

func (handler *Handler) CreateEmptyCollection(c *gin.Context) {
	var req createEmptyRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	collection := &model.Collection{
		UserProfileID: req.UserProfileID,
		Day:           req.Day,
		Name:          req.Name,
	}

	err := handler.CollectionService.CreateEmptyCollection(c, *collection)

	if err != nil {
		log.Panicf("Failed to create empty collection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
