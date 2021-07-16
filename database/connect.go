package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"esystem/model"
	"esystem/conf"
)


var  DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(postgres.Open(conf.DSN),&gorm.Config{})

	// create tables
	db.AutoMigrate(&model.User{}, &model.Goods{}, &model.Log{})


	if err == nil {
		DB = db
		err := db.AutoMigrate(&model.AdminUser{})
		return db, err
	}
	return nil, err
}