package routes

import (
	"location/controller"

	"github.com/labstack/echo"
)

func SubDistrictRoutes(routes *echo.Echo, api controller.SubDistrictController) {

	sub_district := routes.Group("/sub_district")
	{
		sub_district.GET("", api.GetAll)
		sub_district.GET("/:id", api.GetByID)
		sub_district.GET("/district/:district_id", api.GetByDistrictID)
	}
}
