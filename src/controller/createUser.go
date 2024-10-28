package controller

import (
	"github.com/d8barcelos/api-golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("você chamou a rota errada")
	c.JSON(err.Code, err)
}
