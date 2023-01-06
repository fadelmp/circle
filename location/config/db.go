package config

import (
	"location/entity"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {
	processENV()

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

func processENV() {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading env file")
	}
}

func createTable(db *gorm.DB) {
	db.CreateTable(&entity.Country{})
	db.CreateTable(&entity.Province{})
	db.CreateTable(&entity.City{})
	db.CreateTable(&entity.District{})
	db.CreateTable(&entity.SubDistrict{})
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&entity.Country{})
	db.AutoMigrate(&entity.Province{})
	db.AutoMigrate(&entity.City{})
	db.AutoMigrate(&entity.District{})
	db.AutoMigrate(&entity.SubDistrict{})
}
