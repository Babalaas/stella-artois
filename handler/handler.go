package handler

import (
	"babalaas/stella-artois/models"

	"github.com/gin-gonic/gin"
)

// holds required services for handlers to function
type Handler struct {
	PostService models.PostService
}

// holds services injected on handler initilization
type Config struct {
	Router  *gin.Engine
	BaseURL string
}

// can replace parameters with Config parameter object
func NewHandler(router *gin.Engine, postService models.PostService, baseURL string) {
	handler := &Handler{
		PostService: postService,
	}

	postRouteGroup := router.Group("/posts")

	postRouteGroup.GET("/:id", handler.GetPostById)
}
