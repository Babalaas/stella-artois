package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

	router, err := inject()

	if err != nil {
		log.Fatal("Failure injecting data sources")
	}

	log.Println(fmt.Sprintf("Starting Server on port %s", MyEnv.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", MyEnv.Port), router))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", MyEnv.Port),
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}

func inject() (*gin.Engine, error) {
	// repo layer
	postRepo := repository.NewPostRepository()

	// service layer
	postService := service.NewPostService(&postRepo)

	// handler layer
	router := gin.Default()
	router.RedirectTrailingSlash = false
	baseUrl := "/api/v1"

	handler.NewHandler(router, postService, baseUrl)

	return router, nil
}
