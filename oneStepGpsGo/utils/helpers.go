package utils

import (
	"log"
	"time"
)

func ParseTime(timeString string) int64 {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		log.Println("Error parsing date\n[ERROR] -", err.Error())
		return time.Now().Unix()
	}
	return t.Unix()
}
