package main

import (
	"fmt"
	ClientUseCase "study/application/useCase/client"
	"study/config"
	"study/infra/database"
	httpServer "study/infra/httpServer"
	clientRoute "study/infra/httpServer/routes/client"
	"study/infra/repository/client"
)

func main() {
	loadConfig := config.LoadConfig()
	mongoAdapter, err := database.NewMongoDBAdapter(loadConfig.Database.URI, loadConfig.Database.DBName, loadConfig.Database.Username, loadConfig.Database.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	clientRepository := client.NewRepositoryClient(mongoAdapter)
	server := httpServer.NewHTTPServer()
	adapter := httpServer.NewAdapter(server)
	clientUseCase := ClientUseCase.NewUseCaseClient(clientRepository)
	clientRoute.NewClientController(adapter, clientRepository, clientUseCase).Start()

	err = adapter.Listen(":" + loadConfig.HTTP.Port) // Iniciando o servidor
	if err != nil {
		fmt.Println(err)
		return
	}
}
