package controllers

import (
	"net/http"
	"oneStepGps/entities"
	"oneStepGps/models"
	"oneStepGps/services"

	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	var data models.IdModel
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, "Could not read data from the request body.\n[ERROR] - "+err.Error())
	}
	result := services.GetProfileData(data.Id)
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var data entities.User
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Could not read data from the request body", "success": false})
	}

	err = services.UpdateUser(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Could not update data", "success": false})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved changed", "success": true})
}
