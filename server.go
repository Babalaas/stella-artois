package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"babalaas/web-server/db"
	"babalaas/web-server/routes"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter().StrictSlash(true)

	routes.RegisterPostRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", MyEnv.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", MyEnv.Port), router))
}