package controllers

import (
	"babalaas/web-server/db"
	"babalaas/web-server/entities"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entities.Post
	json.NewDecoder(r.Body).Decode(&post)
	db.Instance.Create(&post)
	json.NewEncoder(w).Encode(post)
}
