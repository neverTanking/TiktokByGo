package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = mysqlUrl(Database["user"], Database["password"], Database["db_name"])

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		SkipDefaultTransaction: true, //关闭默认事务
		PrepareStmt:            true, //缓存预编译语句
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{}, &Video{}, &Like{}, &Comment{})
	if err != nil {
		panic(err)
	}
}

func mysqlUrl(user string, pwd string, db_name string) string {
	return user + ":" + pwd + "@tcp(localhost:3306)/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
}
