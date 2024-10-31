package model

import (
	"fmt"

	"github.com/d8barcelos/api-golang/src/configuration/logger"
	"github.com/d8barcelos/api-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println("Senha criptografada:", ud.Password)
	return nil
}
