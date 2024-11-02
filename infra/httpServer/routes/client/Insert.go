package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"study/infra/repository/client"
)

func (c *Controller) Insert() {
	c.Adapter.Handler("/client", "POST", func(w http.ResponseWriter, r *http.Request) {

		_, err := json.Marshal(r.Body)
		if err != nil {
			w.Write([]byte("Erro"))
		}

		b := client.InsertClient{
			Name: "John Doe",
			Age:  30,
			Address: client.Address{
				City:    "New York",
				Zipcode: "10001",
			},
		}

		get, err := c.clientUseCase.InsertOne(b)
		if err != nil {
			w.Write([]byte("Erro"))
		}
		//w.Write()
		w.Write([]byte(fmt.Sprintf("Inserted: %v", get)))
	})

}
