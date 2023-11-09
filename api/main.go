package main

import (
	"github.com/GDGVIT/gibspons-backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// mapping all routes
	controllers.MapRoutes()

	r.Run()
}
