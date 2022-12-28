package controller

import (
	"strconv"

	"github.com/Aman130/HotelManagement/config"
	"github.com/Aman130/HotelManagement/models"
	"github.com/gin-gonic/gin"
)

func BookRoom(c *gin.Context) {
	roomid, _ := strconv.Atoi(c.Query("roomId"))
	custId, _ := strconv.Atoi(c.Query("custId"))

	name := c.Query("name")
	email := c.Query("email")
	contact := c.Query("contact")

	hotel := []models.Hotel{}
	result := config.DB.Where("room_id= ?", roomid).Find(&hotel)

	if result.RowsAffected != 0 {
		c.JSON(200, gin.H{"Message": "Room not present"})
		return
	} else {
		booked := []models.Bookings{}
		result = config.DB.Where("room_id=?", roomid).Find(&booked)
		if result.RowsAffected != 0 {
			c.JSON(200, gin.H{"Message": "Room already taken"})
			return
		} else {
			var customer models.Customer
			customer = models.Customer{
				CustId:  custId,
				Name:    name,
				Contact: contact,
				Email:   email,
			}
			config.DB.Create(&customer)
			var book models.Bookings = models.Bookings{
				RoomId:     roomid,
				CustomerId: custId,
			}
			config.DB.Create(&book)
			c.JSON(200, customer)
		}
	}
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
}
