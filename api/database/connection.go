package database

import (
	"log"

	"github.com/GDGVIT/gibspons-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=" + config.DB.Host +
		" user=" + config.DB.Username +
		" password=" + config.DB.Password +
		" dbname=" + config.DB.Database +
		" port=" + config.DB.Port +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}
	DB = db
	log.Println("[CONNECTION] Database connection established")
}
