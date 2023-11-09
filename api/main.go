package main

import (
	"github.com/GDGVIT/gibspons-backend/config"
	"github.com/GDGVIT/gibspons-backend/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfigs()
}

func main() {
	r := gin.Default()

	// mapping all routes
	controllers.MapRoutes(r)

	r.Run()
}
