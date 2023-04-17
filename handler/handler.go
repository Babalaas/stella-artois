package handler

import (
	"babalaas/stella-artois/model"

	"github.com/gin-gonic/gin"
)

// Handler holds required services for handlers to function
type Handler struct {
	PostService        model.PostService
	UserProfileService model.UserProfileService
	FriendshipService  model.FriendshipService
	CommentService     model.CommentService
}

// Config holds services injected on handler initilization
type Config struct {
	Router  *gin.Engine
	BaseURL string

	PostService        model.PostService
	UserProfileService model.UserProfileService
	FriendshipService  model.FriendshipService
	CommentService     model.CommentService
}

// NewHandler is a factory function which a new Handler struct
func NewHandler(config *Config) {
	handler := &Handler{
		PostService:        config.PostService,
		UserProfileService: config.UserProfileService,
		FriendshipService:  config.FriendshipService,
		CommentService:     config.CommentService,
	}

	postRouteGroup := config.Router.Group("/posts")
	postRouteGroup.GET("/:id", handler.GetPostByID)

	userProfileRouteGroup := config.Router.Group("/user-profiles")
	userProfileRouteGroup.POST("", handler.Register)
	userProfileRouteGroup.POST("/login", handler.LogIn)

	friendshipRouteGroup := config.Router.Group("/friends")
	friendshipRouteGroup.GET("/:id", handler.GetAllFriends)

	commentRouteGroup := config.Router.Group("/post-comments")
	commentRouteGroup.POST("", handler.CreatePostComment)
	commentRouteGroup.DELETE("/:id", handler.DeletePostComment)
}
