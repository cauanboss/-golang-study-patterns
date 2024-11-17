package client

import (
	"net/http"
)

func (c *Controller) Find() {
	c.Adapter.Handler("/client", "GET", func(w http.ResponseWriter, r *http.Request) {
		format := r.URL.Query().Get("format") // Ou r.Header.Get("Accept")
		formatter := c.ChooseFormat(format)

		// Busca os dados do caso de uso
		data, err := c.clientUseCase.Find("")
		if err != nil {
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		// Renderiza a resposta no formato escolhido
		content, contentType, err := formatter.Render(data)
		if err != nil {
			http.Error(w, "Erro ao processar a resposta", http.StatusInternalServerError)
			return
		}

		// Define cabeçalhos e escreve a resposta
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, "Erro ao escrever a resposta", http.StatusInternalServerError)
			return
		}
	})
}
