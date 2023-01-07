// You can edit this code!
// Click here and start typing.
package main

import (
	routes2 "core/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	processENV()
	route := echo.New()
	routes := routes2.Init(route)

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	service_port := os.Getenv("SERVICE_PORT")
	routes.Logger.Fatal(routes.Start(service_port))

}

func processENV() {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading env file")
	}
}
