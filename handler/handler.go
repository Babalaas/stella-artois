package handler

import (
	"babalaas/stella-artois/model"

	"github.com/gin-gonic/gin"
)

// Handler holds required services for handlers to function
type Handler struct {
	PostService        model.PostService
	UserProfileService model.UserProfileService
}

// Config holds services injected on handler initilization
type Config struct {
	Router  *gin.Engine
	BaseURL string

	PostService        model.PostService
	UserProfileService model.UserProfileService
}

// NewHandler is a factory function which a new Handler struct
// with the required services and creates necessary route groups
func NewHandler(config *Config) {
	handler := &Handler{
		PostService:        config.PostService,
		UserProfileService: config.UserProfileService,
	}

	postRouteGroup := config.Router.Group("/posts")
	postRouteGroup.GET("/:id", handler.GetPostByID)

	userProfileRouteGroup := config.Router.Group("/user-profiles")
	userProfileRouteGroup.POST("", handler.Register)
	userProfileRouteGroup.POST("/login", handler.LogIn)
}
