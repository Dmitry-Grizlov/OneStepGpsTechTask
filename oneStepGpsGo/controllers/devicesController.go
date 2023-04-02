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
		c.JSON(http.StatusBadRequest, utils.ErrorBadRequest(err))
		return
	}

	result.Data, err = services.GetDevices(data.ApiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorInternal(err, "Could not get devices."))
		return
	}

	c.JSON(http.StatusOK, result)
}
