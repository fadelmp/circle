package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToBaseDto(entity entity.Base) dto.Base {
	return dto.Base{
		Is_Actived: entity.Is_Actived,
		Is_Deleted: entity.Is_Deleted,
		Created_By: entity.Created_By,
		Updated_By: entity.Updated_By,
		Created_At: entity.Created_At,
		Updated_At: entity.Updated_At,
	}
}
