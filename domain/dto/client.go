package dto

type Address struct {
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type InsertDTO struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}
