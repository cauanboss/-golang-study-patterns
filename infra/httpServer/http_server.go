package httpServer

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	router *mux.Router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{router: mux.NewRouter()}
}

func (h *HTTPServer) Listen(port string) error {
	fmt.Println("HTTPServer Listen")
	return http.ListenAndServe(port, h.router)
}

func (h *HTTPServer) Handler(path string, methods string, handle func(w http.ResponseWriter, r *http.Request)) {
	h.router.HandleFunc(path, handle).Methods(methods)
}
