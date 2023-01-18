package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type StatusRepositoryContract interface {
	GetAll() []entity.Status
	GetByID(uint) entity.Status
	GetByName(string) entity.Status

	Create(entity.Status) error
	Update(entity.Status) error
	Delete(entity.Status) error
	ActiveStatus(entity.Status) error
}

type StatusRepository struct {
	DB *gorm.DB
}

func ProviderStatusRepository(DB *gorm.DB) StatusRepository {
	return StatusRepository{DB: DB}
}

func (s *StatusRepository) GetAll() []entity.Status {

	var statuses []entity.Status

	s.DB.Find(&statuses)

	return statuses
}

func (s *StatusRepository) GetByID(id uint) entity.Status {

	var status entity.Status

	s.DB.Where("id=?", id).Find(&status)

	return status
}

func (s *StatusRepository) GetByName(name string) entity.Status {

	var status entity.Status

	s.DB.Where("name=?", name).Find(&status)

	return status
}

func (s *StatusRepository) Create(status entity.Status) error {

	err := s.DB.Create(&status).Error

	return err
}

func (s *StatusRepository) Update(status entity.Status) error {

	err := s.DB.Model(&status).Update(&status).Error

	return err
}

func (s *StatusRepository) Delete(status entity.Status) error {

	err := s.DB.Model(&status).Where("id=?", status.ID).Updates(map[string]interface{}{
		"is_actived": status.Base.Is_Actived,
		"is_deleted": status.Base.Is_Deleted,
		"updated_at": status.Base.Updated_At,
		"updated_by": status.Base.Updated_By,
	}).Error

	return err
}

func (s *StatusRepository) ActiveStatus(status entity.Status) error {

	err := s.DB.Model(&status).Where("id=?", status.ID).Updates(map[string]interface{}{
		"is_actived": status.Base.Is_Actived,
		"updated_at": status.Base.Updated_At,
		"updated_by": status.Base.Updated_By,
	}).Error

	return err
}
