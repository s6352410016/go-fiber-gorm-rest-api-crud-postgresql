package database

import (
	"fmt"
	"go-rest-api-crud/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("PG_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed To Connected Database\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected To Database Successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migration")
	db.AutoMigrate(&models.Customer{})

	DB = db
}
