package routes

import (
	"location/controller"

	"github.com/labstack/echo"
)

func ProvinceRoutes(routes *echo.Echo, api controller.ProvinceController) {

	province := routes.Group("/province")
	{
		province.GET("", api.GetAll)
		province.GET("/:id", api.GetByID)
	}
}
