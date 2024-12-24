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
	env := config.LoadConfig()
	mongoAdapter, errDB := database.NewMongoDBAdapter(env.Database.URI, env.Database.DBName, env.Database.Username, env.Database.Password)
	if errDB != nil {
		fmt.Println(errDB)
		return
	}
	repositories := repository.NewRepositories(mongoAdapter)
	server := httpServer.NewHTTPServer()
	adapter := httpServer.NewAdapter(server)

	clientUseCase := ClientUseCase.NewUseCaseClient(repositories.Client)
	clientRoute.NewClientController(adapter, clientUseCase).Start()

	errListen := adapter.Listen(":" + env.HTTP.Port)
	if errListen != nil {
		fmt.Println(errListen)
		return
	}
}
