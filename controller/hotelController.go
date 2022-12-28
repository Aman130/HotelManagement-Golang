package controller

import (
	"strconv"
	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/models"
	"github.com/gin-gonic/gin"
)

func GetRoom(c *gin.Context) {
	hotel := []models.Hotel{}
	result := config.DB.Find(&hotel)
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(200, hotel)
}

func QueryRooms(c *gin.Context){
	
	price,_:=strconv.Atoi(c.Query("Price"))
	result :=[]models.Hotel{}
	config.DB.Raw("SELECT * FROM hotels WHERE room_price<= ?", price).Scan(&result)
	c.JSON(200,result)
}

func CreateRoom(c *gin.Context) {
	var hotel models.Hotel
	c.BindJSON(&hotel)
	config.DB.Create(&hotel)
	c.JSON(200, hotel)
}