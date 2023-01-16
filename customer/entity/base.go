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

	var base Base

	base.Is_Actived = true
	base.Is_Deleted = false
	base.Created_By = "System"
	base.Updated_By = "System"
	base.Created_At = time.Now()
	base.Updated_At = time.Now()

	return base
}

func BaseUpdate() Base {

	var base Base

	base.Updated_By = "System"
	base.Updated_At = time.Now()

	return base
}

func BaseDelete() Base {
	var base Base

	base.Is_Actived = false
	base.Is_Deleted = true
	base.Updated_By = "System"
	base.Updated_At = time.Now()

	return base
}

func BaseActivate(is_active bool) Base {

	var base Base

	base.Is_Actived = is_active
	base.Updated_By = "System"
	base.Updated_At = time.Now()

	return base
}
