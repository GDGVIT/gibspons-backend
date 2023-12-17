package database

import (
	"errors"

	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

func CreateSponsorship(sponsorship *tables.Sponsorships) error {
	tx := database.DB.Create(sponsorship)
	return tx.Error
}

func GetSponsorships(pocid int) ([]tables.Sponsorships, error) {
	var spons []tables.Sponsorships
	tx := database.DB.Find(&spons, "poc_id = ?", pocid)
	return spons, tx.Error
}

func UpdateSponsorshipStatus(id int, status string) error {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	spon.Status = status
	tx = database.DB.Save(spon)
	return tx.Error
}

func UpdateSponsorshipEmail(id int, email string) error {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	spon.Email = email
	tx = database.DB.Save(spon)
	return tx.Error
}

func DeleteSponsorship(id int) error {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	if tx.Error != nil {
		return errors.New("sponsorship not found")
	}
	tx = database.DB.Delete(tables.Sponsorships{}, "id = ?", id)
	return tx.Error
}
