package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"babalaas/stella-artois/config"
	"babalaas/stella-artois/handler"
	"babalaas/stella-artois/repository"
	"babalaas/stella-artois/service"

	"github.com/gin-gonic/gin"
)

type envConfig struct {
	Port             string
	ConnectionString string
	GinMode          string
}

func main() {
	myEnv := initEnvConfig()
	config.Connect(myEnv.ConnectionString)
	config.Migrate()

	gin.SetMode(myEnv.GinMode)

	router, err := inject()

	if err != nil {
		log.Fatal("Failure injecting data sources")
	}

	log.Printf("Starting Server on port %s", myEnv.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", myEnv.Port), router))
}

func inject() (*gin.Engine, error) {
	instance := config.GetInstance()
	// repositories
	postRepo := repository.NewPostRepository(instance)
	userProfileRepo := repository.NewUserProfileRepository(instance)
	friendshipRepo := repository.NewFriendshipRepository(instance)
	postCommentRepo := repository.NewCommentRepository(instance)

	// service configs
	userProfileConfig := &service.UPSConfig{
		UserProfileRepository: userProfileRepo,
	}

	friendshipConfig := service.FSConfig{
		FriendshipRepository: friendshipRepo,
	}

	postCommentConfig := &service.CSConfig{
		PostCommentRepo: postCommentRepo,
	}

	// services
	postService := service.NewPostService(postRepo)
	userProfileService := service.NewUserProfileService(userProfileConfig)
	friendshipService := service.NewFriendshipService(friendshipConfig)
	postCommentService := service.NewCommentService(postCommentConfig)

	// handler layer
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.RedirectTrailingSlash = false

	handlerConfig := &handler.Config{
		Router:             router,
		BaseURL:            "",
		PostService:        postService,
		UserProfileService: userProfileService,
		FriendshipService:  friendshipService,
		CommentService:     postCommentService,
	}

	handler.NewHandler(handlerConfig)

	return router, nil
}

func initEnvConfig() *envConfig {
	newEnv := envConfig{
		Port:             os.Getenv("PORT"),
		ConnectionString: os.Getenv("CONNECTION_STRING"),
		GinMode:          os.Getenv("GIN_MODE"),
	}
	return &newEnv
}
