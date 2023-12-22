package controllers

import (
	"net/http"

	database "github.com/GDGVIT/gibspons-backend/database/dbops"
	"github.com/GDGVIT/gibspons-backend/database/tables"
	"github.com/GDGVIT/gibspons-backend/utility"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Phone    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		utility.GinCtxError(c, "Failed to create user, check user body")
		return
	}

	var hashPass string
	var err error
	if hashPass, err = utility.HashPassword(body.Password); err != nil {
		utility.GinCtxError(c, "Failed to hash password")
		return
	}

	user := tables.Users{
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
		Password: []byte(hashPass),
	}

	if err := database.CreateUser(&user); err != nil {
		utility.GinCtxError(c, "Failed to create user, check db logs")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
