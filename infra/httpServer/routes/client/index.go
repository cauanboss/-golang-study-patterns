package client

import (
	ClientUseCase "study/application/useCase/client"
	"study/infra/httpServer"
	"study/infra/repository/client"
)

type Controller struct {
	Adapter       *httpServer.Adapter
	clientUseCase *ClientUseCase.UseCase
}

func NewClientController(adapter *httpServer.Adapter, repository *client.IndexClient, clientUseCase *ClientUseCase.UseCase) *Controller {
	return &Controller{Adapter: adapter, clientUseCase: clientUseCase}
}

func (c *Controller) Start() {
	c.Find()
	c.Insert()
}
