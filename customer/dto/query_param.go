package dto

type QueryParam struct {
	Status string `query:"status"`
	Filter string `query:"filter"`
	Name   string `query:"name"`
}
