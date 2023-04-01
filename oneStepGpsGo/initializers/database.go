package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDbConnection() (*gorm.DB, error) {
	connsectionString := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(connsectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
