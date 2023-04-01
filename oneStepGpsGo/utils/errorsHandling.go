package utils

import (
	"log"
	"oneStepGps/models"
)

func ErrorBadRequest(err error) models.ApiResponseModel {
	log.Println(err.Error())
	return models.ApiResponseModel{Message: "Could not read data from the request body."}
}

func ErrorInternal(err error, msg string) models.ApiResponseModel {
	log.Println(err.Error())
	return models.ApiResponseModel{Message: msg}
}

func ErrorCreateDbConnection(err error) {
	log.Println("Unexpected error occured while creating database connection.\n[ERROR] -", err.Error())
}
