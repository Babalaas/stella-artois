package routes

import (
	"babalaas/web-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.Engine) {
	router.GET("/api/posts", controllers.GetPosts)
	router.GET("/api/posts/:id", controllers.GetPostById)
	router.POST("/api/posts", controllers.CreatePost)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)
}
