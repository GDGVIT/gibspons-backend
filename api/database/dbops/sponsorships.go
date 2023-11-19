package database

import (
	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func CreateSponsorship(sponsorship *tables.Sponsorship) error {
	tx := database.DB.Create(sponsorship)
	return tx.Error
}
