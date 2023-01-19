package dto

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
