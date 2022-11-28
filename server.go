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

type EnvConfig struct {
	Port             string `mapstructure:"PORT"`
	ConnectionString string `mapstructure:"CONNECTION_STRING"`
}

// Viper Code
var MyEnvConfigs *EnvConfig

func InitEnvConfigs() {
	MyEnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *EnvConfig) {
	// viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {

	InitEnvConfigs()
	db.Connect(MyEnvConfigs.ConnectionString)
	db.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	routes.RegisterPostRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", MyEnvConfigs.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", MyEnvConfigs.Port), router))

}
