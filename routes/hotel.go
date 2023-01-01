package routes

import (
	// "bytes"
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"


	"github.com/Aman130/HotelManagement/controller"
	// "github.com/Aman130/HotelManagement/models"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

// func guidMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		var obj models.Hotel
// 		// obj1 := models.Hotel{
// 		// 	RoomNo:      10,
// 		// 	FloorNo:     10,
// 		// 	RoomService: "Yes",
// 		// 	RoomPrice:   5000,
// 		// }
// 		// fmt.Println(obj1)
// 		if errA := c.ShouldBind(&obj); errA == nil {
// 			// c.String(http.StatusOK, `the body should be formA`)
// 			fmt.Println("Hello", obj)
// 			obj.RoomId="10"

// 			b,_:=json.Marshal(obj)


// 			// body, _ := ioutil.ReadAll(c.Request.Body)
// 			// println(string(body))

// 			c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))
// 			// Always an error is occurred by this because c.Request.Body is EOF now.
// 		}
// 		// uuid:= uuid.New()
// 		c.Next()

// 	}
// }
func HotelRoutes(router *gin.Engine) {

	hotelapi:= router.Group("/api/hotel")
	{
		hotelapi.GET("/getAllRooms", controller.GetRoom)
		hotelapi.GET("/queryRooms", controller.QueryRooms)
		hotelapi.POST("/createRoom", controller.CreateRoom)
	}

	customerapi:= router.Group("/api/customer")
	{
		customerapi.GET("/getInfo",controller.GetInfo)
		customerapi.GET("/fetchCustomer",controller.FetchInfo)
		customerapi.POST("/addCustomer",controller.AddCustomer)
	}

	bookingapi:= router.Group("/api/booking")
	{
		bookingapi.POST("/bookRoom",controller.BookRoom)
	}
}
