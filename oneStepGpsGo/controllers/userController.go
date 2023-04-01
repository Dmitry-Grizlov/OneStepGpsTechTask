package controllers

import (
	"log"
	"net/http"
	"oneStepGps/entities"
	"oneStepGps/models"
	"oneStepGps/services"
	"oneStepGps/utils"

	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	var data models.IdModel
	var result models.ApiResponseModel

	err := c.Bind(&data)
	if err != nil {
		result = models.ApiResponseModel{Message: "Could not read data from the request body"}
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result.Data, err = services.GetProfileData(data.Id)
	if err != nil {
		if err != nil {
			log.Println(err.Error())
			result = models.ApiResponseModel{Message: "Could not get user data."}
			c.JSON(http.StatusInternalServerError, result)
			return
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var data entities.User
	var result models.ApiResponseModel

	err := c.Bind(&data)
	if err != nil {
		result = utils.ErrorBadRequest(err)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = services.UpdateUser(data)
	if err != nil {
		result = utils.ErrorInternal(err, "Could not update user")
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
