package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /{id}
func (handler *Handler) GetPostById(c *gin.Context) {
	reqId := c.Param("id")

	/*
		if !exists {
			log.Panicf("Unable to extract post from request context for unknown reason: %v\n", c)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Post record not found!"})
			return
		}
	*/

	uid := uuid.Must(uuid.Parse(reqId))
	ctx := c.Request.Context()
	post, err := handler.PostService.GetById(ctx, uid)

	if err != nil {
		log.Panicf("Unable to find post")
		c.JSON(http.StatusNotFound, gin.H{"error": "Post record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
