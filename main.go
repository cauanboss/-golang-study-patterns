package main

import (
	"fmt"
	ClientUseCase "study/application/useCase/client"
	"study/config"
	"study/infra/database"
	httpServer "study/infra/httpServer"
	clientRoute "study/infra/httpServer/routes/client"
	"study/infra/repository"
)

func main() {
	config.LoadConfig()
	mongoAdapter, err := database.NewMongoDBAdapter(config.Env.Database.URI, config.Env.Database.DBName, config.Env.Database.Username, config.Env.Database.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	repositories := repository.NewRepositories(mongoAdapter)
	server := httpServer.NewHTTPServer()
	adapter := httpServer.NewAdapter(server)

	clientUseCase := ClientUseCase.NewUseCaseClient(repositories.Client)
	clientRoute.NewClientController(adapter, repositories.Client, clientUseCase).Start()

	err = adapter.Listen(":" + config.Env.HTTP.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
