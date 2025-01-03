package client

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (c *Controller) Find() {
	c.Adapter.Handler("/client", "GET", func(w http.ResponseWriter, r *http.Request) {
		format := r.URL.Query().Get("format") // Ou r.Header.Get("Accept")
		formatter := c.ChooseFormat(format)
		data, err := c.clientUseCase.Find("")
		if err != nil {
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}
		content, contentType, err := formatter.Render(data)
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

	c.Adapter.Handler("/client/{id}", "GET", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		format := r.URL.Query().Get("format")
		formatter := c.ChooseFormat(format)
		data, err := c.clientUseCase.Find(id)
		if err != nil {
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}
		content, contentType, err := formatter.Render(data)
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
