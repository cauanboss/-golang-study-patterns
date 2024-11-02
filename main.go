package main

import (
	"fmt"
	ClientUseCase "study/application/useCase/client"
	"study/infra/database"
	httpServer "study/infra/httpServer"
	clientRoute "study/infra/httpServer/routes/client"
	"study/infra/repository/client"
)

func main() {
	mongoAdapter, err := database.NewMongoDBAdapter("mongodb://localhost:27017/study", "study")
	if err != nil {
		fmt.Println(err)
		return
	}
	clientRepository := client.NewRepositoryClient(mongoAdapter)
	server := httpServer.NewHTTPServer()
	adapter := httpServer.NewAdapter(server)
	clientUseCase := ClientUseCase.NewUseCaseClient(clientRepository)
	clientRoute.NewClientController(adapter, clientRepository, clientUseCase).Start() // Criando uma nova instância de Routes

	err = adapter.Listen(":8080") // Iniciando o servidor
	if err != nil {
		fmt.Println(err)
		return
	}
}
