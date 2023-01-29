package injection

import (
	"order/controller"
	repository "order/repository"
	"order/usecase"

	"github.com/jinzhu/gorm"
)

func ArticleInjection(db *gorm.DB) controller.ArticleController {

	ArticleRepository := repository.ProviderArticleRepository(db)
	ArticleUsecase := usecase.ProviderArticleUsecase(ArticleRepository)
	ArticleController := controller.ProviderArticleController(ArticleUsecase)

	return ArticleController
}
