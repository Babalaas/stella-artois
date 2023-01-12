package db

import (
	"babalaas/stella-artois/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var instance *gorm.DB
var err error

func Connect(connectionString string) {

	instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	instance.AutoMigrate(&models.Post{})
	log.Println("Database Migration Completed...")
}

func GetInstance() *gorm.DB {
	return instance
}
