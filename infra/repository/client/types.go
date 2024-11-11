package client

type Address struct {
	City    string
	Zipcode string
}

// TODO: Change the struct name to InsertClient
type InsertClient struct {
	Name    string
	Age     int
	Address Address
}
