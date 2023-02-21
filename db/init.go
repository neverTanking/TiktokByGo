package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DSN = mysqlUrl(Database["user"], Database["password"], Database["db_name"])

var DB *gorm.DB

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{}, &Video{}, &Like{}, &Comment{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}

}

func mysqlUrl(user string, pwd string, db_name string) string {
	return user + ":" + pwd + "@tcp(localhost:3306)/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
}
