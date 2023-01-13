package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}
