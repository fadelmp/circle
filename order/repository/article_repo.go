package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type ArticleRepositoryContract interface {
	Create(entity.Article) (entity.Article, error)
}

type ArticleRepository struct {
	DB *gorm.DB
}

func ProviderArticleRepository(DB *gorm.DB) ArticleRepository {
	return ArticleRepository{DB: DB}
}

func (a *ArticleRepository) Create(article entity.Article) (entity.Article, error) {

	err := a.DB.Create(&article).Error

	return article, err
}
