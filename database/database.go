package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	// TODO: Add migrations
	// db.AutoMigrate(&models.Answer{}, &models.Company{}, &models.Document{}, &models.Event{}, &models.Question{}, &models.Series{}, &models.Tag{}, &models.User{})

	Database = DbInstance{Db: db}
}
