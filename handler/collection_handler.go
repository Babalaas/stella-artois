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

type getUserCollectionsRequest struct {
	UserProfileID uuid.UUID `json:"user_profile_id" binding:"required"`
}

// CreateEmptyCollection is the HTTP handler for a user_profile to create an empty collection
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

	c.JSON(http.StatusCreated, gin.H{"success": "created"})
}

// DeleteCollection is HTTP handler to delete one collection by id
func (handler *Handler) DeleteCollection(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	err := handler.CollectionService.Delete(ctx, uid)

	if err != nil {
		log.Println("Could not delete")
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": "successfully deleted collection",
	})
}

func (handler *Handler) GetUserCollections(c *gin.Context) {
	var req getUserCollectionsRequest

	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		log.Panicf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	collections, err := handler.CollectionService.GetUserCollections(c, req.UserProfileID)

	if err != nil {
		log.Printf("Could not get user's collections: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"collections": collections})
}
