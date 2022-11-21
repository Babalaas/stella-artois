package routes

import (
	"babalaas/web-server/controllers"

	"github.com/gorilla/mux"
)

func RegisterPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/post/", controllers.CreatePost).Methods("POST")
}
