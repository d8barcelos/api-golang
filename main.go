package main

import (
	"log"

	"github.com/d8barcelos/api-golang/src/controller"
	"github.com/d8barcelos/api-golang/src/controller/routes"
	"github.com/d8barcelos/api-golang/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// mongodb.NewMongoDbInitConnection()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
