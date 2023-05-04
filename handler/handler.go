package handler

import (
	"babalaas/stella-artois/model"
	"babalaas/stella-artois/service"

	"github.com/gin-gonic/gin"
)

// Handler holds required services for handlers to function
type Handler struct {
	PostService        model.PostService
	UserProfileService model.UserProfileService
	FriendshipService  model.FriendshipService
	CommentService     model.CommentService
	ReactionService    model.ReactionService
	FeedService        service.FeedService
	CollectionService  model.CollectionService
}

// Config holds services injected on handler initilization
type Config struct {
	Router  *gin.Engine
	BaseURL string

	PostService        model.PostService
	UserProfileService model.UserProfileService
	FriendshipService  model.FriendshipService
	CommentService     model.CommentService
	ReactionService    model.ReactionService
	FeedService        service.FeedService
	CollectionService  model.CollectionService
}

// NewHandler is a factory function which a new Handler struct
func NewHandler(config *Config) {
	handler := &Handler{
		PostService:        config.PostService,
		UserProfileService: config.UserProfileService,
		FriendshipService:  config.FriendshipService,
		CommentService:     config.CommentService,
		ReactionService:    config.ReactionService,
		FeedService:        config.FeedService,
		CollectionService:  config.CollectionService,
	}

	postRouteGroup := config.Router.Group("/posts")
	postRouteGroup.GET("/:id", handler.GetPostByID)
	postRouteGroup.POST("", handler.UploadPost)

	userProfileRouteGroup := config.Router.Group("/user-profiles")
	userProfileRouteGroup.POST("", handler.Register)
	userProfileRouteGroup.POST("/login", handler.LogIn)
	userProfileRouteGroup.GET("", handler.Search)
	userProfileRouteGroup.GET("/:id", handler.GetUserProfileByID)

	friendshipRouteGroup := config.Router.Group("/friends")
	friendshipRouteGroup.GET("/:id", handler.GetAllFriends)
	friendshipRouteGroup.GET("/requests/:id", handler.GetFriendRequests)
	friendshipRouteGroup.GET("", handler.SearchNonFriends)
	friendshipRouteGroup.POST("", handler.RequestFriend)
	friendshipRouteGroup.PATCH("", handler.AcceptFriend)
	friendshipRouteGroup.DELETE("", handler.RemoveFriend)

	commentRouteGroup := config.Router.Group("/post-comments")
	commentRouteGroup.POST("", handler.CreatePostComment)
	commentRouteGroup.DELETE("/:id", handler.DeletePostComment)
	commentRouteGroup.GET("/:id", handler.GetAllComments)

	reactionRouteGroup := config.Router.Group("/post-reactions")
	reactionRouteGroup.POST("", handler.ReactToPost)

	feedRouteGroup := config.Router.Group("/feed")
	feedRouteGroup.GET("/:id", handler.GenerateFeed)
	feedRouteGroup.GET("/collectons/:id", handler.GetPostsInCollection)

	collectionRouteGroup := config.Router.Group("/collections")
	collectionRouteGroup.POST("", handler.CreateEmptyCollection)
	collectionRouteGroup.DELETE("/:id", handler.DeleteCollection)
	collectionRouteGroup.GET("", handler.GetUserCollections)
	collectionRouteGroup.PATCH("", handler.UpdateCollection)
	collectionRouteGroup.POST("/posts", handler.AddPostToCollection)
}
