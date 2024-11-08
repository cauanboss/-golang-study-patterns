package domain

type Address struct {
	City    string
	Zipcode string
}

type Client struct {
	Name    string
	Age     int
	Address Address
}
