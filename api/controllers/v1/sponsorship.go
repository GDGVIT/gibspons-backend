package controllers

import (
	"net/http"
	"strconv"

	database "github.com/GDGVIT/gibspons-backend/database/dbops"
	"github.com/GDGVIT/gibspons-backend/database/tables"
	"github.com/GDGVIT/gibspons-backend/utility"
	"github.com/gin-gonic/gin"
)

func CreateSponsorship(c *gin.Context) {
	var body struct {
		PocID       int
		CompanyName string
		Email       string
		Description string
		Status      string
	}

	if err := c.Bind(&body); err != nil {
		utility.GinCtxError(c, "Failed to create sponsorship, check body")
		return
	}

	spon := tables.Sponsorships{
		PocID:        body.PocID,
		CompanyName:  body.CompanyName,
		CurrentEmail: body.Email,
		Description:  body.Description,
		Status:       body.Status,
	}

	if err := database.CreateSponsorship(&spon); err != nil {
		utility.GinCtxError(c, "Failed to create sponsorship, check db logs")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"sponsorship": spon,
	})
}

func GetSponsorshipById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	spon, err := database.GetSponsorshipById(idInt)
	if err != nil {
		utility.GinCtxError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sponsorship": spon,
	})
}

func GetSponsorshipsByCompanyName(c *gin.Context) {
	name := c.Query("name")
	spons, err := database.GetSponsorshipsByCompanyName(name)
	if err != nil {
		utility.GinCtxError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sponsorships": spons,
	})
}

func GetSponsorships(c *gin.Context) {
	pocid, exists := c.GetQuery("pocid")
	if !exists {
		utility.GinCtxError(c, "Please pass in POC ID")
		return
	}

	pocidInt, _ := strconv.Atoi(pocid)
	spons, err := database.GetSponsorships(pocidInt)
	if err != nil {
		utility.GinCtxError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sponsorships": spons,
	})
}

func UpdateSponsorshipStatus(c *gin.Context) {
	id, exists := c.GetQuery("id")
	if !exists {
		utility.GinCtxError(c, "Please pass in Sponsorship ID")
		return
	}
	status, stat_exists := c.GetQuery("status")
	email, email_exists := c.GetQuery("email")
	if !(stat_exists || email_exists) {
		utility.GinCtxError(c, "Please pass in status")
		return
	}

	idInt, _ := strconv.Atoi(id)
	if stat_exists {
		if err := database.UpdateSponsorshipStatus(idInt, status); err != nil {
			utility.GinCtxError(c, "Failed to update sponsorship, check db logs")
			return

		}
	}
	if email_exists {
		if err := database.UpdateSponsorshipEmail(idInt, email); err != nil {
			utility.GinCtxError(c, "Failed to update sponsorship, check db logs")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated sponsorship",
	})
}

func DeleteSponsorship(c *gin.Context) {
	id, exists := c.GetQuery("id")
	if !exists {
		utility.GinCtxError(c, "Please pass in sponsorship ID")
		return
	}

	idInt, _ := strconv.Atoi(id)
	if err := database.DeleteSponsorship(idInt); err != nil {
		utility.GinCtxError(c, "Failed to delete sponsorship, check db logs")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted sponsorship",
	})
}
