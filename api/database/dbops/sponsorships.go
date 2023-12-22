package database

import (
	"errors"

	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/GDGVIT/gibspons-backend/database/tables"
)

// create new sponsorship
func CreateSponsorship(sponsorship *tables.Sponsorships) error {
	tx := database.DB.Create(sponsorship)
	return tx.Error
}

// get poc's sponsorships
func GetSponsorships(pocid int) ([]tables.Sponsorships, error) {
	var spons []tables.Sponsorships
	tx := database.DB.Find(&spons, "poc_id = ?", pocid)
	return spons, tx.Error
}

// get spon by id
func GetSponsorshipById(id int) (tables.Sponsorships, error) {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	return spon, tx.Error
}

// get all spons associated with company name
func GetSponsorshipsByCompanyName(name string) ([]tables.Sponsorships, error) {
	var spons []tables.Sponsorships
	tx := database.DB.Find(&spons, "company_name = ?", name)
	return spons, tx.Error
}

// update spon status
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

// update spon email
func UpdateSponsorshipEmail(id int, email string) error {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}

	spon.CurrentEmail = email
	tx = database.DB.Save(spon)
	return tx.Error
}

// delete spon
func DeleteSponsorship(id int) error {
	var spon tables.Sponsorships
	tx := database.DB.First(&spon, "id = ?", id)
	if tx.Error != nil {
		return errors.New("sponsorship not found")
	}
	tx = database.DB.Delete(tables.Sponsorships{}, "id = ?", id)
	return tx.Error
}
