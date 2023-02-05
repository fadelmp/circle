package repository

import (
	"strconv"
	"user/config"
	entity "user/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type PrivilegeRepositoryContract interface {
	GetAll() []entity.Privilege
	GetByID(uint) entity.Privilege
	GetByName(string) entity.Privilege

	Create(entity.Privilege) error
	Update(entity.Privilege) error
	ActiveStatus(entity.Privilege) error
}

type PrivilegeRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderPrivilegeRepository(DB *gorm.DB, Redis *redis.Client) PrivilegeRepository {
	return PrivilegeRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (p *PrivilegeRepository) GetAll() []entity.Privilege {

	var privileges []entity.Privilege

	query := p.DB.Find(&privileges)
	keys := "privileges"

	// Get Privilege All
	config.CheckRedisQuery(p.Redis, query, keys)

	return privileges
}

func (p *PrivilegeRepository) GetByID(id uint) entity.Privilege {

	var privilege entity.Privilege

	query := p.DB.Where("id=?", id).Find(&privilege)
	keys := "privilege_id_" + strconv.FormatUint(uint64(id), 10)

	// Get Privilege By Id
	config.CheckRedisQuery(p.Redis, query, keys)

	return privilege
}

func (p *PrivilegeRepository) GetByName(name string) entity.Privilege {

	var privilege entity.Privilege

	query := p.DB.Where("name=?", name).Find(&privilege)
	keys := "privilege_name_" + name

	// Get Privilege by Name
	config.CheckRedisQuery(p.Redis, query, keys)

	return privilege
}

func (p *PrivilegeRepository) Create(privilege entity.Privilege) error {

	// Create Privilege
	return p.DB.Create(&privilege).Error
}

func (p *PrivilegeRepository) Update(privilege entity.Privilege) error {

	// update Privilege by id
	return p.DB.Model(&privilege).Update(&privilege).Error
}

func (p *PrivilegeRepository) ChangeStatus(privilege entity.Privilege) error {

	// delete Privilege by id, by change is active value to false
	return p.DB.Model(&privilege).Where("id=?", privilege.ID).Updates(map[string]interface{}{
		"is_actived": privilege.Base.Is_Actived,
		"is_deleted": privilege.Base.Is_Deleted,
		"updated_at": privilege.Base.Updated_At,
		"updated_by": privilege.Base.Updated_By,
	}).Error
}
