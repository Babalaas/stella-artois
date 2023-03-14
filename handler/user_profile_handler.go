package handler

import (
	"babalaas/stella-artois/model"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	resID, registerErr := handler.UserProfileService.Register(c.Request.Context(), reqUserProfile)

	if registerErr != nil {
		log.Panicf("Failed to register user profile: %v\n", registerErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": registerErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": resID,
	})
}
