package httpServer

import "net/http"

type AdapterInterface interface {
	Listen(port string) error
	Handler(path string, methods string, handle func(w http.ResponseWriter, r *http.Request))
}
