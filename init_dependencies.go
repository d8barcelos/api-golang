package main

import (
	"github.com/d8barcelos/api-golang/src/controller"
	"github.com/d8barcelos/api-golang/src/model/repository"
	"github.com/d8barcelos/api-golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}