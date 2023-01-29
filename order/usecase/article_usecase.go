package usecase

import (
	"order/dto"
	"order/mapper"
	repository "order/repository"
)

type ArticleUsecaseContract interface {
	GetAll() []dto.Article
	GetByID(uint) dto.Article
}

type ArticleUsecase struct {
	ArticleRepository repository.ArticleRepository
}

func ProviderArticleUsecase(a repository.ArticleRepository) ArticleUsecase {
	return ArticleUsecase{
		ArticleRepository: a,
	}
}

func (a *ArticleUsecase) GetAll() []dto.Article {

	articles := a.ArticleRepository.GetAll()

	return mapper.ToArticleDtoList(articles)
}

func (a *ArticleUsecase) GetByID(id uint) dto.Article {

	article := a.ArticleRepository.GetByID(id)

	return mapper.ToArticleDto(article)
}
