package controllers

import (
	"babalaas/stella-artois/db"
	"babalaas/stella-artois/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePostInput struct {
	User_ID       string `json:"user_id" binding:"required"`
	Collection_ID string `json:"collection_id"`
	Caption       string `json:"caption" binding:"required"`
	Location      string `json:"location" binding:"required"`
	Image         string `json:"image" binding:"required"`
	Image2        string `json:"image2"`
	Drink_Number  int    `json:"drink_number" binding:"required"`
}

type UpdatePostInput struct {
	Collection_ID string `json:"collection_id"`
	Caption       string `json:"caption" binding:"required"`
	Location      string `json:"location" binding:"required"`
	Drink_Number  int    `json:"drink_number" binding:"required"`
}

// GET /posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	db.GetInstance().Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GET /posts/{id}
func GetPostById(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := db.GetInstance().Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// POST /posts
func CreatePost(c *gin.Context) {
	// Validate input
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	post := models.Post{
		ID:            "",
		Profile_ID:    input.User_ID,
		Collection_ID: uuid.Nil.String(),
		Caption:       input.Caption,
		Location:      input.Location,
		Created:       time.Time{},
		Image:         input.Image,
		Image2:        input.Image2,
		Drink_Number:  input.Drink_Number,
		Like_Count:    0,
	}

	db.GetInstance().Create(&post)

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// PUT /posts/{id}
func UpdatePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := db.GetInstance().Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post record not found!"})
		return
	}

	// Validate input
	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newInput := models.Post{
		Collection_ID: input.Collection_ID,
		Caption:       input.Caption,
		Location:      input.Location,
		Drink_Number:  input.Drink_Number,
	}

	db.GetInstance().Model(&post).Updates(newInput)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/{id}
func DeletePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := db.GetInstance().Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.GetInstance().Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
