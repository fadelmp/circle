package routes

import (
	"location/controller"

	"github.com/labstack/echo"
)

func DistrictRoutes(routes *echo.Echo, api controller.DistrictController) {

	district := routes.Group("/District")
	{
		district.GET("", api.GetAll)
		district.GET("/:id", api.GetByID)
		district.GET("/city/:city_id", api.GetByCityID)
	}
}
