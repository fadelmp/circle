package usecase

import (
	"order/dto"
	"order/mapper"
	repository "order/repository"
)

type ArticleUsecaseContract interface {
	Check(uint, []dto.Article) error
	Create(dto.Article) error
}

type ArticleUsecase struct {
	ArticleRepository repository.ArticleRepository
	ServiceUsecase    ServiceUsecase
}

func ProviderArticleUsecase(
	a repository.ArticleRepository,
	s ServiceUsecase,
) ArticleUsecase {
	return ArticleUsecase{
		ArticleRepository: a,
		ServiceUsecase:    s,
	}
}

func (a *ArticleUsecase) Check(order_id uint, articles []dto.Article) error {

	for _, value := range articles {

		value.OrderID = order_id

		if err := a.Create(value); err != nil {
			return err
		}
	}

	return nil
}

func (a *ArticleUsecase) Create(dto dto.Article) error {

	article_entity := mapper.ToArticleEntity(dto)

	article, err := a.ArticleRepository.Create(article_entity)
	if err != nil {
		return err
	}

	if err = a.ServiceUsecase.Check(article.ID, dto.Services); err != nil {
		return err
	}

	return nil
}
