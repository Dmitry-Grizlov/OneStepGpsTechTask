package main

import (
	"oneStepGps/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(route *gin.Engine) {
	route.POST("/login", controllers.Login)
	route.POST("/devices", controllers.DevicesList)
	route.POST("/user", controllers.UserProfile)
	route.POST("/updateUser", controllers.UpdateUser)
}
