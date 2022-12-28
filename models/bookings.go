package models

type Bookings struct {
	RoomId     int `json:"roomId" gorm:"primaryKey"`
	CustomerId int `json:"customerId"`
}
