package httpServer

import (
	"fmt"
	"net/http"
)

type Adapter struct {
	server AdapterInterface
}

func NewAdapter(server AdapterInterface) *Adapter {
	return &Adapter{server: server}
}

func (adapter *Adapter) Listen(port string) error {
	fmt.Printf("Starting server on port %s\n", port)
	if err := adapter.server.Listen(port); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
		return err
	}
	return nil
}

func (adapter *Adapter) Handler(path string, methods string, handle func(w http.ResponseWriter, r *http.Request)) {
	fmt.Println("Adapter Handler", path, methods)
	adapter.server.Handler(path, methods, handle)
}
