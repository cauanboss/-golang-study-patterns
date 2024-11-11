package client

import (
	ClientUseCase "study/application/useCase/client"
	"study/infra/httpServer"
)

type Controller struct {
	Adapter       *httpServer.Adapter
	clientUseCase *ClientUseCase.UseCase
}

func NewClientController(adapter *httpServer.Adapter, clientUseCase *ClientUseCase.UseCase) *Controller {
	return &Controller{Adapter: adapter, clientUseCase: clientUseCase}
}

func (c *Controller) Start() {
	c.Find()
	c.Insert()
}
