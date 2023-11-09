package database

import (
	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func CreateUser(user *tables.Users) error {
	tx := database.DB.Create(user)
	return tx.Error
}
