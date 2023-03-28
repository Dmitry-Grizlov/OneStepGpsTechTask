package services

import (
	"encoding/json"
	"log"
	"net/http"
	"oneStepGps/models"
	"oneStepGps/utils"
	"os"
)

func GetDevices(key string) []models.DeviceModel {
	url := "/device-info?state_detail=1&api-key="
	var response []models.DeviceModel
	sendRequest("GET", url, key, &response)
	for i := 0; i < len(response); i++ {
		response[i].TrackerTime = utils.ParseTime(response[i].DtTracker)
		response[i].StatusTime = utils.ParseTime(response[i].DriveStatusBeginTime)
	}
	return response
}

func sendRequest(method, url, key string, target interface{}) any {
	url = os.Getenv("ONE_STEP_GPS_URL") + url + key
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error creating request.\n[ERROR] -", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
