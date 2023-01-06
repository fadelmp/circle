package routes

import (
	"location/controller"

	"github.com/labstack/echo"
)

func CountryRoutes(routes *echo.Echo, api controller.CountryController) {

	country := routes.Group("/country")
	{
		country.GET("", api.GetAll)
		country.GET("/:id", api.GetByID)
	}
}
