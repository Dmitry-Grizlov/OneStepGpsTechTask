package models

type LoginFilterModel struct {
	ApiKey string `json:"apiKey,omitempty"`
}

type LoginModel struct {
	ApiKey string `json:"apiKey,omitempty"`
	Id     int    `json:"id,omitempty"`
}
