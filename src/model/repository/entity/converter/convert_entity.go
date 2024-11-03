package converter

import (
	"github.com/d8barcelos/api-golang/src/model"
	"github.com/d8barcelos/api-golang/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.Id.Hex())
	return domain
}
