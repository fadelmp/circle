package controller

import (
	"order/config"
	"order/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	ArticleUsecase usecase.ArticleUsecase
}

func ProviderArticleController(a usecase.ArticleUsecase) ArticleController {
	return ArticleController{
		ArticleUsecase: a,
	}
}

func (a *ArticleController) GetAll(e echo.Context) error {

	articles := a.ArticleUsecase.GetAll()

	if len(articles) == 0 {
		return SuccessResponse(e, nil, config.ArticleNotFound)
	}

	return SuccessResponse(e, articles, config.GetArticleSuccess)
}

func (a *ArticleController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	article := a.ArticleUsecase.GetByID(uint(id))

	if article.ID == 0 {
		return SuccessResponse(e, nil, config.ArticleNotFound)
	}

	return SuccessResponse(e, article, config.GetArticleSuccess)
}
