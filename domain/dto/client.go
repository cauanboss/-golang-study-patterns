package dto

type AddressDTO struct {
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type InsertDTO struct {
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Address AddressDTO `json:"address"`
}

type FindDTO struct {
	ID string `json:"id"`
}

type FindManyDTO struct {
	ID []string `json:"id"`
}
