package config

import (
	"order/entity"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", db_username+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_name+"?parseTime=true")
	if err != nil {
		logrus.Error("Cannot Connect to MySQL DB")
		panic(err)
	}

	createTable(db)
	migrateDDL(db)
	return db
}

func createTable(db *gorm.DB) {
	db.CreateTable(&entity.Status{})
	db.CreateTable(&entity.Unit{})
	db.CreateTable(&entity.Order{})
	db.CreateTable(&entity.Article{})
	db.CreateTable(&entity.Service{})
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&entity.Status{})
	db.CreateTable(&entity.Unit{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.Article{})
	db.AutoMigrate(&entity.Service{})
}
