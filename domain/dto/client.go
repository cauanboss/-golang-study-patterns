package dto

type AddressDTO struct {
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type InsertDTO struct {
	Address AddressDTO `json:"address" xml:"address"`
	Age     int        `json:"age" xml:"age"`
	Name    string     `json:"name" xml:"name"`
}

type FindDTO struct {
	ID string `json:"id" xml:"id"`
}

type FindManyDTO struct {
	ID []string `json:"id" xml:"id"`
}
