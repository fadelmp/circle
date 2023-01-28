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
		Unit:       dto.Unit,
		Amount:     dto.Amount,
		Note:       dto.Note,
		Path_Image: dto.Path_Image,
		Services:   ToServiceEntityList(dto.Services),
	}
}

func ToArticleEntityList(dto []dto.Article) []entity.Article {

	articles := make([]entity.Article, len(dto))

	for i, value := range dto {
		articles[i] = ToArticleEntity(value)
	}

	return articles

}
