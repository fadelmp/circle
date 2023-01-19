package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ReadEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		godotenv.Load("/home/oshwinwa/circle-be/core/.env")
		logrus.Error("Error loading env file")
	}
}
