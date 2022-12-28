package main

import (
	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	
	routes.HotelRoutes(router)

	router.Run(":8080")
}