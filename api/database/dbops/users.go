package database

import (
	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func CreateUser(user *tables.Users) error {
	tx := database.DB.Create(user)
	return tx.Error
}

func FetchUser(email string) *tables.Users {
	var user *tables.Users
	database.DB.Where("user_email = ?", email).First(&user)
	return user
}
