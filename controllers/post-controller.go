package controllers

import (
	"babalaas/web-server/db"
	"babalaas/web-server/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET /api/posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []entities.Post
	db.Instance.Find(&posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// GET /api/posts/{id}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfPostExists(productId) == false {
		json.NewEncoder(w).Encode("Post Not Found!")
		return
	}
	var post entities.Post
	db.Instance.First(&post, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// POST /api/posts
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entities.Post
	json.NewDecoder(r.Body).Decode(&post)
	db.Instance.Create(&post)
	json.NewEncoder(w).Encode(post)
}

// PUT api/posts/{id}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]

	if checkIfPostExists(postId) == false {
		json.NewEncoder(w).Encode("Post Not Found!")
		return
	}

	var post entities.Post
	db.Instance.First(&post, postId)
	json.NewDecoder(r.Body).Decode(&post)
	db.Instance.Save(&post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// DELETE /api/posts/{id}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postId := mux.Vars(r)["id"]
	if checkIfPostExists(postId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Post Not Found!")
		return
	}
	var post entities.Post
	db.Instance.Delete(&post, postId)
	json.NewEncoder(w).Encode("Post Deleted Successfully!")
}

func checkIfPostExists(postId string) bool {
	var post entities.Post
	db.Instance.First(&post, postId)
	if post.ID == 0 {
		return false
	}
	return true
}
