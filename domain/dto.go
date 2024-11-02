package domain

type UseCaseClientFound struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address struct {
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
}
