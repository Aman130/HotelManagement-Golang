package controller

import (
	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/models"
	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	var user models.Customer
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func FetchInfo(c *gin.Context){
	name:= c.Query("Name")
	result :=[]models.Customer{}
	config.DB.Raw("SELECT * FROM customers WHERE name = ?", name).Scan(&result)
	c.JSON(200,result)
}

func AddCustomer(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	config.DB.Create(&customer)
	c.JSON(200, customer)
}