package config

import (
	"babalaas/stella-artois/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var instance *gorm.DB
var err error

// Connect opens a session with a PostgreSQL database using the passed connection string
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

// Migrate keeps model up to date with database instance
func Migrate() {
	instance.AutoMigrate(&model.Post{})
	instance.AutoMigrate(&model.UserProfile{})
	log.Println("Database Migration Completed...")
}

// GetInstance returns current database instance
func GetInstance() *gorm.DB {
	return instance
}
