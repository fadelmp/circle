package routes

import (
	"core/controller"

	"github.com/labstack/echo"
)

func OrderArticleRoutes(routes *echo.Echo, api controller.OrderArticleController) {

	article := routes.Group("/order/article")
	{
		article.GET("", api.GetOrderArticles)
		article.GET("/:id", api.GetOrderArticleById)
	}

}
