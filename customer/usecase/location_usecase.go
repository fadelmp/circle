package usecase

import (
	"customer/request"
)

type LocationUsecaseContract interface {
}

type LocationUsecase struct {
	LocationRequest request.LocationRequest
}

func ProviderLocationUsecase(l request.LocationRequest) LocationUsecase {
	return LocationUsecase{
		LocationRequest: l,
	}
}

// Implementation
