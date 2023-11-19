package controllers

import (
	"net/http"

	database "github.com/GDGVIT/gibspons-backend/database/dbops"
	"github.com/GDGVIT/gibspons-backend/database/tables"
	"github.com/GDGVIT/gibspons-backend/utility"
	"github.com/gin-gonic/gin"
)

func CreateSponsorship(c *gin.Context) {
	var body struct {
		POCName     string
		Email       string
		Description string
		Status      string
	}

	if err := c.Bind(&body); err != nil {
		utility.GinCtxError(c, "Failed to create sponsorship, check body")
		return
	}

	spon := tables.Sponsorship{
		POCName:     body.POCName,
		Email:       body.Email,
		Description: body.Description,
		Status:      body.Status,
	}

	if err := database.CreateSponsorship(&spon); err != nil {
		utility.GinCtxError(c, "Failed to create sponsorship, check db logs")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"sponsorship": spon,
	})
}
