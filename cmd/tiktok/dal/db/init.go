package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN = "root:@tcp(localhost:3306)/tiktok?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{}, &Video{}, &Like{}, &Comment{})
	if err != nil {
		panic(err)
	}
}
