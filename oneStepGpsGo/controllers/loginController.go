package controllers

import (
	"net/http"
	"oneStepGps/models"
	"oneStepGps/services"
	"oneStepGps/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.LoginFilterModel
	var result models.ApiResponseModel

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBadRequest(err))
		return
	}

	result.Data, err = services.Login(data.ApiKey)
	if err != nil {
		result.Message = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
