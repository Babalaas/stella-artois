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

type registerRequest struct {
	DisplayName string    `json:"display_name" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Birthdate   time.Time `json:"birthdate" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	ProfilePic  string    `json:"profile_pic" binding:"required"`
}

type logInRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type searchRequest struct {
	Query string `json:"query" binding:"required"`
}

// Register handles the HTTP request to create one new user_profile entity
// and store it in the database.
func (handler *Handler) Register(c *gin.Context) {
	var req registerRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind user profile JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	reqUserProfile := &model.UserProfile{
		DisplayName: req.DisplayName,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		Birthdate:   req.Birthdate,
		Password:    req.Password,
		ProfilePic:  req.ProfilePic,
	}

	resUserProfile, registerErr := handler.UserProfileService.Register(c.Request.Context(), reqUserProfile)

	if registerErr != nil {
		log.Panicf("Failed to register user profile: %v\n", registerErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": registerErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_profile": resUserProfile,
	})
}

// LogIn authenticates one user_profile
func (handler *Handler) LogIn(c *gin.Context) {
	var req logInRequest

	if bindErr := c.ShouldBind(&req); bindErr != nil {
		log.Panicf("Failed to bind login JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	u := &model.UserProfile{
		DisplayName: req.DisplayName,
		Password:    req.Password,
	}

	ctx := c.Request.Context()
	userProfile, err := handler.UserProfileService.LogIn(ctx, u)

	if err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_profile": userProfile,
	})
}

// Search returns a list of users wiith displayanmes that contain the query text
func (handler *Handler) Search(c *gin.Context) {
	var req searchRequest

	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		log.Panicf("Failed to bind search JSON input: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", bindErr)})
		return
	}

	query := req.Query
	ctx := c.Request.Context()
	userProfiles, err := handler.UserProfileService.Search(ctx, query)

	if err != nil {
		log.Printf("Failed to seach for user profiles: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": userProfiles,
	})
}

// GetUserProfileByID cleans the id parameter and calls the PostService to get a post by based on the parameter (GET /{id})
func (handler *Handler) GetUserProfileByID(c *gin.Context) {
	reqID := c.Param("id")

	uid := uuid.Must(uuid.Parse(reqID))
	ctx := c.Request.Context()
	userProfile, err := handler.UserProfileService.GetByID(ctx, uid)

	if err != nil {
		log.Println("Unable to find post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           userProfile.ID,
		"display_name": userProfile.DisplayName,
		"first_name":   userProfile.FirstName,
		"last_name":    userProfile.LastName,
		"email":        userProfile.Email,
		"phone":        userProfile.Phone,
		"birthdate":    userProfile.Birthdate,
		"password":     "",
		"profile_pic":  userProfile.ProfilePic,
	})
}
