package controllers

import (
	"log"
	"net/http"
	"oneStepGps/models"
	"oneStepGps/services"

	"github.com/gin-gonic/gin"
)

func DevicesList(c *gin.Context) {
	var data models.LoginModel
	var result models.ApiResponseModel
	err := c.Bind(&data)
	if err != nil {
		log.Println(err.Error())
		result = models.ApiResponseModel{Message: "Could not read data from the request body."}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result.Data, err = services.GetDevices(data.ApiKey)
	if err != nil {
		result = models.ApiResponseModel{Message: "Could not get devices.\n[ERROR] - " + err.Error()}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
