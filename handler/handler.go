package handler

import (
	"babalaas/stella-artois/model"

	"github.com/gin-gonic/gin"
)

// Handler holds required services for handlers to function
type Handler struct {
	PostService model.PostService
}

// Config holds services injected on handler initilization
type Config struct {
	Router  *gin.Engine
	BaseURL string
}

// NewHandler creates a new Handler struct with the required services and creates necessary route groups (can replace parameters with Config parameter object)
func NewHandler(router *gin.Engine, postService model.PostService, baseURL string) {
	handler := &Handler{
		PostService: postService,
	}

	postRouteGroup := router.Group("/posts")

	postRouteGroup.GET("/:id", handler.GetPostByID)
}
