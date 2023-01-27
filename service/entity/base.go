package entity

import (
	"time"
)

type Base struct {
	Is_Actived bool      `json:"is_actived"`
	Is_Deleted bool      `json:"is_deleted"`
	Created_By string    `json:"created_by"`
	Updated_By string    `json:"updated_by"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func BaseCreate() Base {
	return Base{
		Is_Actived: true,
		Is_Deleted: false,
		Created_By: "System",
		Updated_By: "System",
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
}

func BaseUpdate() Base {
	return Base{
		Updated_By: "System",
		Updated_At: time.Now(),
	}
}

func BaseDelete() Base {
	return Base{
		Is_Actived: false,
		Is_Deleted: true,
		Updated_By: "System",
		Updated_At: time.Now(),
	}
}

func BaseActivate(is_actived bool) Base {

	return Base{
		Is_Actived: is_actived,
		Is_Deleted: false,
		Updated_By: "System",
		Updated_At: time.Now(),
	}
}
