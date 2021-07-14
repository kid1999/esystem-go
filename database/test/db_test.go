/**
@auther: kid1999
@date: 2021/7/14 9:24
@desc: db_test
**/

package test

import (
	config "esystem/conf"
	"esystem/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestDatabase(t *testing.T) {
	conf := config.Get()
	db, _ := gorm.Open(postgres.Open(conf.DSN),&gorm.Config{})


	db.Create(&model.AdminUser{})
	db.Create(&model.User{})
}
