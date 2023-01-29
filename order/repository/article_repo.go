package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type ArticleRepositoryContract interface {
	GetAll() []entity.Article
	GetByID(uint) entity.Article
}

type ArticleRepository struct {
	DB *gorm.DB
}

func ProviderArticleRepository(DB *gorm.DB) ArticleRepository {
	return ArticleRepository{DB: DB}
}

func (a *ArticleRepository) GetAll() []entity.Article {

	var articles []entity.Article

	a.DB.Model(&entity.Article{}).Order("id asc").
		Preload("Order").Preload("Status").Preload("Service").
		Find(&articles)

	return articles
}

func (a *ArticleRepository) GetByID(id uint) entity.Article {

	var article entity.Article

	a.DB.Where("id=?", id).Order("id asc").
		Preload("Order").Preload("Status").Preload("Service").
		Find(&article)

	return article
}
