package controllers

import (
	"net/http"
	"oneStepGps/models"
	"oneStepGps/services"
	"oneStepGps/utils"

	"github.com/gin-gonic/gin"
)

func DevicesList(c *gin.Context) {
	var data models.LoginFilterModel
	var result models.ApiResponseModel

	err := c.Bind(&data)
	if err != nil {
		result = utils.ErrorBadRequest(err)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result.Data, err = services.GetDevices(data.ApiKey)
	if err != nil {
		result = utils.ErrorInternal(err, "Could not get devices.")
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
