package main

import (
	"babalaas/web-server/db"
	"babalaas/web-server/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type envConfigs struct {
	Port             string `mapstructure:"PORT"`
	ConnectionString string `mapstructure:"CONNECTION_STRING"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()

	return
}

func main() {

	InitEnvConfigs()
	db.Connect(EnvConfigs.ConnectionString)
	db.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	routes.RegisterPostRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", EnvConfigs.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", EnvConfigs.Port), router))

}
