package repository

import (
	"strconv"
	"user/config"
	entity "user/entity"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type UserRepositoryContract interface {
	GetAll() []entity.User
	GetByID(uint) entity.User
	GetByName(string) entity.User

	Create(entity.User) error
	Update(entity.User) error
	Delete(entity.User) error
	ActiveStatus(entity.User) error
}

type UserRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func ProviderUserRepository(DB *gorm.DB, Redis *redis.Client) UserRepository {
	return UserRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (u *UserRepository) GetAll() []entity.User {

	var users []entity.User

	query := u.DB.Find(&users)
	keys := "users"

	// Get User All
	config.CheckRedisQuery(u.Redis, query, keys)

	return users
}

func (u *UserRepository) GetByID(id uint) entity.User {

	var user entity.User

	query := u.DB.Where("id=?", id).Find(&user)
	keys := "user_id_" + strconv.FormatUint(uint64(id), 10)

	// Get User By Id
	config.CheckRedisQuery(u.Redis, query, keys)

	return user
}

func (u *UserRepository) GetByName(name string) entity.User {

	var user entity.User

	query := u.DB.Where("name=?", name).Find(&user)
	keys := "user_name_" + name

	// Get User by Name
	config.CheckRedisQuery(u.Redis, query, keys)

	return user
}

func (u *UserRepository) Create(user entity.User) error {

	// Create User
	err := u.DB.Create(&user).Error

	return err
}

func (u *UserRepository) Update(user entity.User) error {

	// update User by id
	err := u.DB.Model(&user).Update(&user).Error

	return err
}

func (u *UserRepository) Delete(user entity.User) error {

	// delete User by id, by change is active value to false
	err := u.DB.Model(&user).Where("id=?", user.ID).Updates(map[string]interface{}{
		"is_actived": user.Base.Is_Actived,
		"is_deleted": user.Base.Is_Deleted,
		"updated_at": user.Base.Updated_At,
		"updated_by": user.Base.Updated_By,
	}).Error

	return err
}

func (u *UserRepository) ActiveStatus(user entity.User) error {

	// delete User by id, by change is active value to false
	err := u.DB.Model(&user).Where("id=?", user.ID).Updates(map[string]interface{}{
		"is_actived": user.Base.Is_Actived,
		"updated_at": user.Base.Updated_At,
		"updated_by": user.Base.Updated_By,
	}).Error

	return err
}
