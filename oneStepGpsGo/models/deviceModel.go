package models

type DeviceModel struct {
	Id                   int
	DeviceId             string  `json:"device_id"`
	DisplayName          string  `json:"display_name"`
	Make                 string  `json:"make"`
	Model                string  `json:"model"`
	ActiveState          string  `json:"active_state"`
	Lat                  float64 `json:"lat"`
	Lng                  float64 `json:"lng"`
	DtTracker            string  `json:"dt_tracker"`
	DriveStatus          string  `json:"drive_status"`
	DriveStatusBeginTime string  `json:"drive_status_begin_time"`
	SpeedMph             float64 `json:"speed_mph"`
	SpeedKph             float64 `json:"speed_kph"`
	AltitudeM            float64 `json:"altitude_m"`
	Vin                  string  `json:"vin"`
	OdometerMi           float64 `json:"odometer_mi"`
	OdometerKm           float64 `json:"odometer_km"`
	TrackerTime          int64   `json:"tracker_time"`
	StatusTime           int64   `json:"status_time"`
}
