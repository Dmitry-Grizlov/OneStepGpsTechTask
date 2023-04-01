package services

import (
	"encoding/json"
	"log"
	"net/http"
	"oneStepGps/models"
	"oneStepGps/utils"
	"os"
)

func GetDevices(key string) ([]models.DeviceModel, error) {
	url := "/device-info?state_detail=1&api-key="
	var response []models.DeviceModel

	err := sendRequest("GET", url, key, &response)
	if err != nil {
		log.Println("Error sending request.\n[ERROR] -", err.Error())
		return response, err
	}

	for i := 0; i < len(response); i++ {
		response[i].TrackerTime = utils.ParseTime(response[i].DtTracker)
		response[i].StatusTime = utils.ParseTime(response[i].DriveStatusBeginTime)
	}

	return response, nil
}

func pingServer(apiKey string) []models.DeviceModel {
	pingResult, err := GetDevices(apiKey)
	if err != nil {
		return []models.DeviceModel{}
	}

	return pingResult
}

func sendRequest(method, url, key string, target interface{}) error {
	url = os.Getenv("ONE_STEP_GPS_URL") + url + key

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error creating request.\n[ERROR] -", err.Error())
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Println("Error decoding response.\n[ERROR] -", err)
		return err
	}

	return nil
}
