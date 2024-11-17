package client

import (
	ClientUseCase "study/application/useCase/client"
	"study/infra/httpServer"
	"study/infra/presenter"
)

type Controller struct {
	Adapter       *httpServer.Adapter
	clientUseCase *ClientUseCase.UseCase
	ChooseFormat  func(format string) presenter.Formatter
}

func NewClientController(adapter *httpServer.Adapter, clientUseCase *ClientUseCase.UseCase) *Controller {
	return &Controller{Adapter: adapter, clientUseCase: clientUseCase, ChooseFormat: presenter.Choose}
}

func (c *Controller) Start() {
	c.Find()
	c.Insert()
}
