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

	district := routes.Group("/district")
	{
		district.GET("", api.GetAllDistrict)
		district.GET("/:id", api.GetDistrictByID)
		district.GET("/city/:city_id", api.GetDistrictByCityID)
	}

	sub_district := routes.Group("/sub_district")
	{
		sub_district.GET("", api.GetAllSubDistrict)
		sub_district.GET("/:id", api.GetSubDistrictByID)
		sub_district.GET("/district/:district_id", api.GetSubDistrictByDistrictID)
	}
}
