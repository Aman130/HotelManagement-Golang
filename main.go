package main

import (
	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/controller"
	"github.com/Aman130/HotelManagement/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	controller.CacheIntialize("localhost:6379",0,100)
	routes.HotelRoutes(router)

	router.Run(":8080")
}