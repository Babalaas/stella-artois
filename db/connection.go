package db

import (
	"babalaas/web-server/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Post{})
	log.Println("Database Migration Completed...")
}
