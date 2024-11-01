package controller

import (
	"net/http"

	"github.com/d8barcelos/api-golang/src/configuration/logger"
	"github.com/d8barcelos/api-golang/src/configuration/validation"
	"github.com/d8barcelos/api-golang/src/controller/model/request"
	"github.com/d8barcelos/api-golang/src/model"
	"github.com/d8barcelos/api-golang/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
