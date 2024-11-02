package httpServer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	server := NewHTTPServer()

	// Cria um manipulador simples para a rota "/"
	server.Handler("/", "GET", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Cria um servidor de teste
	testServer := httptest.NewServer(server.router) // Use o router do servidor
	defer testServer.Close()

	// Faz uma requisição GET para o servidor de teste
	resp, err := http.Get(testServer.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	// Verifica o status da resposta
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verifica o corpo da resposta
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	resp.Body.Close()

	assert.Equal(t, "Hello, world!", string(body[:n]))
}
