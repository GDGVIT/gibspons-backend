package controllers

import (
	"github.com/GDGVIT/gibspons-backend/controllers/v1"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	// <------------- USER AUTH ROUTES --------------->
	userGroup := v1.Group("/users")

	userGroup.POST("/signup", controllers.Signup)
}
