package repository

import (
	"strconv"
	"user/config"
	entity "user/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type RoleRepositoryContract interface {
	GetAll() []entity.Role
	GetByID(uint) entity.Role
	GetByName(string) entity.Role

	Create(entity.Role) error
	Update(entity.Role) error
	ActiveStatus(entity.Role) error
}

type RoleRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderRoleRepository(DB *gorm.DB, Redis *redis.Client) RoleRepository {
	return RoleRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *RoleRepository) GetAll() []entity.Role {

	var roles []entity.Role

	query := r.DB.Find(&roles)
	keys := "roles"

	// Get Role All
	config.CheckRedisQuery(r.Redis, query, keys)

	return roles
}

func (r *RoleRepository) GetByID(id uint) entity.Role {

	var role entity.Role

	query := r.DB.Where("id=?", id).Find(&role)
	keys := "role_id_" + strconv.FormatUint(uint64(id), 10)

	// Get Role By Id
	config.CheckRedisQuery(r.Redis, query, keys)

	return role
}

func (r *RoleRepository) GetByName(name string) entity.Role {

	var role entity.Role

	query := r.DB.Where("name=?", name).Find(&role)
	keys := "role_name_" + name

	// Get Role by Name
	config.CheckRedisQuery(r.Redis, query, keys)

	return role
}

func (r *RoleRepository) Create(role entity.Role) error {

	// Create Role
	return r.DB.Create(&role).Error
}

func (r *RoleRepository) Update(role entity.Role) error {

	// update Role by id
	return r.DB.Model(&role).Update(&role).Error
}

func (r *RoleRepository) ActiveStatus(role entity.Role) error {

	// delete Role by id, by change is active value to false
	return r.DB.Model(&role).Where("id=?", role.ID).Updates(map[string]interface{}{
		"is_actived": role.Base.Is_Actived,
		"is_deleted": role.Base.Is_Deleted,
		"updated_at": role.Base.Updated_At,
		"updated_by": role.Base.Updated_By,
	}).Error
}
