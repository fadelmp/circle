package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ReadEnv() {

	prod_err := godotenv.Load("/home/oshwinwa/circle-be/order/.env")

	if prod_err != nil {

		local_err := godotenv.Load(".env")
		if local_err != nil {
			logrus.Error("Error loading env local file")
		}

	}
}
