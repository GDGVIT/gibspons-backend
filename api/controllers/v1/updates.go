package controllers

import (
	"net/http"
	"strconv"
	"time"

	database "github.com/GDGVIT/gibspons-backend/database/dbops"
	"github.com/GDGVIT/gibspons-backend/database/tables"
	"github.com/GDGVIT/gibspons-backend/utility"
	"github.com/gin-gonic/gin"
)

func AddUpdate(c *gin.Context) {
	var body struct {
		ID     int
		Status string
		Email  string
	}

	if err := c.Bind(&body); err != nil {
		utility.GinCtxError(c, "Failed to add update, check body")
		return
	}

	update := tables.Updates{
		SponsorshipsID: uint(body.ID),
		Email:          body.Email,
		Status:         body.Status,
		UpdatedAt:      time.Now(),
	}

	if err := database.AddUpdate(&update); err != nil {
		utility.GinCtxError(c, "Failed to add update, check db logs: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"update": update,
	})
}

func GetUpdates(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	updates, err := database.GetUpdates(idInt)
	if err != nil {
		utility.GinCtxError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updates": updates,
	})
}
