package client

import (
	"encoding/json"
	"net/http"
	"study/domain/dto"
)

func (c *Controller) Insert() {
	c.Adapter.Handler("/client", "POST", func(w http.ResponseWriter, r *http.Request) {

		format := r.URL.Query().Get("format") // Ou r.Header.Get("Accept")
		formatter := c.ChooseFormat(format)

		_, err := json.Marshal(r.Body)
		if err != nil {
			http.Error(w, "Erro ao processar a resposta", http.StatusInternalServerError)
			return
		}

		b := dto.InsertDTO{
			Name: "John Doe",
			Age:  30,
			Address: dto.AddressDTO{
				City:    "New York",
				Zipcode: "10001",
			},
		}

		get, err := c.clientUseCase.InsertOne(b)

		content, contentType, err := formatter.Render(get)
		if err != nil {
			http.Error(w, "Erro ao processar a resposta", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, "Erro ao escrever a resposta", http.StatusInternalServerError)
			return
		}
	})

}
