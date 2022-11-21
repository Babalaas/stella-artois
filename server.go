package main

import (
	"babalaas/web-server/db"
	"babalaas/web-server/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()

	db.Connect(AppConfig.ConnectionString)
	db.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	routes.RegisterPostRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
