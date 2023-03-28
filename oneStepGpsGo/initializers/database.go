package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	connsectionString := os.Getenv("CONNECTION_STRING")
	DB, err = gorm.Open(mysql.Open(connsectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database\n[ERROR] -", err.Error())
	}
}
