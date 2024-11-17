package domain

type Address struct {
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type Client struct {
	Address Address `json:"address" xml:"address"`
	Age     int     `json:"age" xml:"age"`
	Name    string  `json:"name" xml:"name"`
}
