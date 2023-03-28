package controllers

import (
	"net/http"
	"oneStepGps/models"
	"oneStepGps/services"

	"github.com/gin-gonic/gin"
)

func DevicesList(c *gin.Context) {
	var data models.LoginModel
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, "Could not read data from the request body.\n[ERROR] - "+err.Error())
	}
	c.JSON(http.StatusOK, services.GetDevices(data.ApiKey))
}
