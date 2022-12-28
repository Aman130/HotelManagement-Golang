package models


type Hotel struct {
	RoomId      int    `json:"roomId" gorm:"primaryKey"`
	RoomNo      int    `json:"roomNo" gorm:"autoIncrement:true"`
	FloorNo     int    `json:"floorNo"`
	RoomService string `json:"roomService"`
	RoomPrice   int    `json:"roomPrice"`
}
