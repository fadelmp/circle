package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func LocationRoutes(routes *echo.Echo, api controller.LocationController) {

	country := routes.Group("/country")
	{
		country.GET("", api.GetAllCountry)
		country.GET("/:id", api.GetCountryByID)
	}

	province := routes.Group("/province")
	{
		province.GET("", api.GetAllProvince)
		province.GET("/:id", api.GetProvinceByID)
		province.GET("/country/:country_id", api.GetProvinceByCountryID)
	}

	city := routes.Group("/city")
	{
		city.GET("", api.GetAllCity)
		city.GET("/:id", api.GetCityByID)
		city.GET("/province/:province_id", api.GetCityByProvinceID)
	}
}
