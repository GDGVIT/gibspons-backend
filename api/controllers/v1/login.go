package controllers

import (
	"net/http"

	database "github.com/GDGVIT/gibspons-backend/database/dbops"
	"github.com/GDGVIT/gibspons-backend/utility"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		utility.GinCtxError(c, "Failed to create user, check user body")
		return
	}

	user := database.FetchUser(body.Email)
	if user.ID == 0 {
		utility.GinCtxError(c, "Invalid Email")
		return
	}

	verified, err := utility.VerifyPassword(user.Password, body.Password)

	if !verified || err != nil {
		utility.GinCtxError(c, "Invalid Password")
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"user": user,
	})

}
