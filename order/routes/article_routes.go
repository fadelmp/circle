package routes

import (
	"order/controller"

	"github.com/labstack/echo"
)

func ArticleRoutes(routes *echo.Echo, api controller.ArticleController) {

	article := routes.Group("/order/article")
	{
		article.GET("", api.GetAll)
		article.GET("/:id", api.GetByID)
	}
}
