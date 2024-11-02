package client

import (
	"net/http"
)

func (c *Controller) Find() {
	c.Adapter.Handler("/client", "GET", func(w http.ResponseWriter, r *http.Request) {
		get, err := c.clientUseCase.Find("")
		if err != nil {
			w.Write([]byte("Erro"))
		}
		w.Write(get)
	})
}
