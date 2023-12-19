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
	userGroup.POST("/login", controllers.Login)

	// <----------- SPONSORSHIP ROUTES -------------->
	sponGroup := v1.Group("/sponsorships")

	sponGroup.POST("/create", controllers.CreateSponsorship)
	sponGroup.GET("/", controllers.GetSponsorships)
	sponGroup.GET("/", controllers.GetSponsorshipById)
	sponGroup.GET("/update", controllers.UpdateSponsorshipStatus)
	sponGroup.DELETE("/", controllers.DeleteSponsorship)

}
