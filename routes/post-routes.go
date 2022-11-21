package routes

import (
	"babalaas/web-server/controllers"

	"github.com/gorilla/mux"
)

func RegisterPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")
}
