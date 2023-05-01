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

type updateCollectionRequest struct {
	ID            uuid.UUID `json:"id"`
	UserProfileID uuid.UUID `json:"user_profile_id"`
	Day           time.Time `json:"day"`
	Name          string    `json:"name"`
}

type addPostToCollectionRequest struct {
	CollectionID uuid.UUID `json:"collection_id"`
	PostID       uuid.UUID `json:"post_id"`
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

// GetUserCollections is the HTTP handler for return all collections with the same user profile id
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

// UpdateCollection updates the time and day of a collection
func (handler *Handler) UpdateCollection(c *gin.Context) {
	var req updateCollectionRequest

	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		log.Printf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.CollectionService.UpdateCollection(c, model.Collection(req))

	if err != nil {
		log.Printf("Could not get user's collections: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// AddPostToCollection creates a collection post entity
func (handler *Handler) AddPostToCollection(c *gin.Context) {
	var req addPostToCollectionRequest

	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		log.Printf("Failed to bind post comment JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	err := handler.CollectionService.AddPostToCollection(c, req.PostID, req.CollectionID)

	if err != nil {
		log.Printf("Could not add post to collection: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
