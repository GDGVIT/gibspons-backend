package database

import (
	"log"

	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func AutoMigrate() {
	var users tables.Users

	db := DB

	if err := db.AutoMigrate(&users); err != nil {
		log.Fatal("[MIGRATOR] Users Migration Failed")
	}

	log.Println("[MIGRATOR] Migrated all tables")
}
