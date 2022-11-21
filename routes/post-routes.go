package routes

import (
	"babalaas/web-server/controllers"

	"github.com/gorilla/mux"
)

func RegisterPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts/", controllers.CreatePost).Methods("POST")
}