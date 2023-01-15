package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"babalaas/stella-artois/db"
	"babalaas/stella-artois/handler"
	"babalaas/stella-artois/repository"
	"babalaas/stella-artois/service"

	"github.com/gin-gonic/gin"
)

type Env struct {
	Port             string
	ConnectionString string
}

func initEnv() *Env {

	NewEnv := Env{
		Port:             os.Getenv("PORT"),
		ConnectionString: os.Getenv("CONNECTION_STRING"),
	}

	return &NewEnv
}

func main() {
	MyEnv := initEnv()
	db.Connect(MyEnv.ConnectionString)
	db.Migrate()

	gin.SetMode(os.Getenv("GIN_MODE"))

	router, err := inject()

	if err != nil {
		log.Fatal("Failure injecting data sources")
	}

	log.Printf("Starting Server on port %s", MyEnv.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", MyEnv.Port), router))
}

func inject() (*gin.Engine, error) {
	// repo layer
	postRepo := repository.NewPostRepository()

	// service layer
	postService := service.NewPostService(&postRepo)

	// handler layer
	router := gin.Default()
	//router.Use(middleware.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.RedirectTrailingSlash = false
	//baseUrl := "/api/v1"

	handler.NewHandler(router, postService, "/")

	return router, nil
}
