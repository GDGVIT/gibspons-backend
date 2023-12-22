package database

import (
	"log"

	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func AutoMigrate() {
	var users tables.Users
	var sponsorships tables.Sponsorships
	var updates tables.Updates

	db := DB

	if err := db.AutoMigrate(&users); err != nil {
		log.Fatal("[MIGRATOR] Users Migration Failed")
	}
	if err := db.AutoMigrate(&updates); err != nil {
		log.Fatal("[MIGRATOR] Updates Migration Failed")
	}

	if err := db.AutoMigrate(&sponsorships); err != nil {
		log.Fatal("[MIGRATOR] Sponsorships Migration Failed")
	}

	log.Println("[MIGRATOR] Migrated all tables")
}
