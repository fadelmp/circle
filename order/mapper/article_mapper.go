package mapper

import (
	"order/dto"
	"order/entity"
)

func ToArticleEntity(dto dto.Article) entity.Article {
	return entity.Article{
		ID:         dto.ID,
		Name:       dto.Name,
		Quantity:   dto.Quantity,
		UnitID:     dto.UnitID,
		Amount:     dto.Amount,
		Note:       dto.Note,
		Path_Image: dto.Path_Image,
		Services:   ToServiceEntityList(dto.Services),
	}
}

func ToArticleDto(entity entity.Article) dto.Article {
	return dto.Article{
		ID:         entity.ID,
		OrderID:    entity.OrderID,
		Name:       entity.Name,
		Quantity:   entity.Quantity,
		UnitName:   entity.Unit.Name,
		Amount:     entity.Amount,
		Note:       entity.Note,
		Path_Image: entity.Path_Image,
		StatusName: entity.Status.Name,
		Services:   ToServiceDtoList(entity.Services),
	}
}

func ToArticleEntityList(dto []dto.Article) []entity.Article {

	articles := make([]entity.Article, len(dto))

	for i, value := range dto {
		articles[i] = ToArticleEntity(value)
	}

	return articles

}

func ToArticleDtoList(entity []entity.Article) []dto.Article {

	articles := make([]dto.Article, len(entity))

	for i, value := range entity {
		articles[i] = ToArticleDto(value)
	}

	return articles
}
