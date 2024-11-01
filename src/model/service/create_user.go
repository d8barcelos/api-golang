package service

import (
	"fmt"

	"github.com/d8barcelos/api-golang/src/configuration/logger"
	"github.com/d8barcelos/api-golang/src/configuration/rest_err"
	"github.com/d8barcelos/api-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(ud)
	return nil
}
