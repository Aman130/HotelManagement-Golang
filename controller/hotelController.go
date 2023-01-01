package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Aman130/HotelManagement/cache"
	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/models"
	"github.com/gin-gonic/gin"
)

var roomCache cache.RedisCache

func CacheIntialize(host string, db int, exp time.Duration) {
	roomCache = *cache.NewRedisCache(host, db, exp)
}
func GetRoom(c *gin.Context) {

	hotel := []*models.Hotel{}
	result := config.DB.Find(&hotel)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, hotel)
}

func QueryRooms(c *gin.Context) {
	var finder string = "QueryRooms" + c.Query("Price")
	var room []models.Hotel = roomCache.Get(finder)
	if room != nil {
		fmt.Println("hello")
		c.JSON(200, room)
		return
	}
	price, _ := strconv.Atoi(c.Query("Price"))
	result := []models.Hotel{}
	config.DB.Raw("SELECT * FROM hotels WHERE room_price<= ?", price).Scan(&result)
	roomCache.Set(finder, result)
	fmt.Println("Query")
	c.JSON(200, result)
}

func CreateRoom(c *gin.Context) {
	var hotel models.Hotel
	c.BindJSON(&hotel)
	config.DB.Create(&hotel)
	c.JSON(200, hotel)
}
