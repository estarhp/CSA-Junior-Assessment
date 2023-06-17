package utils

import (
	"github.com/google/uuid"
	"time"
)

func FormatTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func GenerateID() string {
	id := uuid.New().String()

	return id
}
