package client

type Address struct {
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

// TODO: Change the struct name to InsertClient
type InsertClient struct {
	Address Address `json:"address" xml:"address"`
	Age     int     `json:"age" xml:"age"`
	Name    string  `json:"name" xml:"name"`
}
