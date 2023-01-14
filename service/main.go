// You can edit this code!
// Click here and start typing.
package main

import (
	"os"
	config "service/config"
	routes2 "service/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	route := echo.New()
	// route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// }))

	dbConfig := config.InitDB()
	routes := routes2.Init(route, dbConfig)

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
