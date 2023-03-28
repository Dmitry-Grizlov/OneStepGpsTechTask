package controllers

import (
	"log"
	"net/http"
	"oneStepGps/models"
	"oneStepGps/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.LoginModel
	err := c.Bind(&data)
	if err != nil {
		log.Fatal("Could not read data from the request body.\n[ERROR] -", err)
	}

	res, err := services.Login(data.ApiKey)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"apiKey": data.ApiKey, "id": res})
}
