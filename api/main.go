package main

import (
	"github.com/GDGVIT/gibspons-backend/config"
	"github.com/GDGVIT/gibspons-backend/controllers"
	"github.com/GDGVIT/gibspons-backend/database"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfigs()
}

func main() {
	_ = database.Connection()

	r := gin.Default()

	// mapping all routes
	controllers.MapRoutes(r)

	r.Run()
}
