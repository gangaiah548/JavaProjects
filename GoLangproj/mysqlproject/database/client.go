package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/entities"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}
