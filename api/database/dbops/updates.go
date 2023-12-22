package database

import (
	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func AddUpdate(u *tables.Updates) error {
	tx := database.DB.Create(u)
	return tx.Error
}

func GetUpdates(id int) ([]tables.Updates, error) {
	var updates []tables.Updates
	tx := database.DB.Find(&updates, "sponsorships_id = ?", id)
	return updates, tx.Error
}
