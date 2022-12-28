package models

type Customer struct {
	CustId  int    `json:"custId" gorm:"primaryKey"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Email   string `json:"mail"`
}
