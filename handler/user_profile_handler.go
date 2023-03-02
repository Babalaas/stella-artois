package handler

import (
	"babalaas/stella-artois/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	DisplayName string    `json:"display_name" binding:"required,display_name"`
	FirstName   string    `json:"first_name" binding:"required,first_name"`
	LastName    string    `json:"last_name" binding:"required,last_name"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone" binding:"required,phone"`
	Gender      string    `json:"gender" binding:"gender"`
	Birthdate   time.Time `json:"birthdate" binding:"required,birthdate"`
	Password    string    `json:"password" binding:"required,password"`
}

func (handler *Handler) Register(c *gin.Context) {
	var req registerRequest

	req_user_profile := &model.UserProfile{
		DisplayName: req.DisplayName,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		Gender:      req.Gender,
		Birthdate:   req.Birthdate,
		Password:    req.Password,
	}

	registerErr := handler.UserProfileService.Register(c.Request.Context(), req_user_profile)

	if registerErr != nil {
		log.Panicf("Failed to register user profile: %v\n", registerErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": registerErr,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"created": "created",
	})
}
