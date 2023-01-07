package routes

import (
	"location/controller"

	"github.com/labstack/echo"
)

func CityRoutes(routes *echo.Echo, api controller.CityController) {

	city := routes.Group("/city")
	{
		city.GET("", api.GetAll)
		city.GET("/:id", api.GetByID)
		city.GET("/province/:province_id", api.GetByProvinceID)
	}
}
