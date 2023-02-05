package usecase

import (
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
)

type UserUsecaseContract interface {
	Create(entity.User) error
	Update(entity.User) error
}

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func ProviderUserUsecase(c repository.UserRepository) UserUsecase {
	return UserUsecase{UserRepository: c}
}

// Implementation

func (c *UserUsecase) Create(dto dto.User, customer_id uint) error {

	// Map dto to entity
	User_entity := mapper.ToUserEntity(dto, customer_id)

	// Create User Data
	err := c.UserRepository.Create(User_entity)

	// Map entity to dto
	return err
}

func (c *UserUsecase) Update(dto dto.User, customer_id uint) error {

	// Map dto to entity
	User_entity := mapper.ToUserEntity(dto, customer_id)

	// Update User Data
	err := c.UserRepository.Update(User_entity)

	// Map entity to DTO
	return err
}
