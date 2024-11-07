package repository

import (
	"context"
	"os"

	"github.com/d8barcelos/api-golang/src/configuration/logger"
	"github.com/d8barcelos/api-golang/src/configuration/rest_err"
	"github.com/d8barcelos/api-golang/src/model"
	"github.com/d8barcelos/api-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init create user repository")

	collectionName := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
