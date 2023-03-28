package entities

type User struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Email  string
	ApiKey string
}
