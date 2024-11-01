package view

import (
	"github.com/d8barcelos/api-golang/src/controller/model/response"
	"github.com/d8barcelos/api-golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		Id:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
