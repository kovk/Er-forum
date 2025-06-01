package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"newgo/global"
	"time"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	Db, err := db.DB()
	if err != nil {
		log.Fatalf("database config error:%v,err")
	}
	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxLifetime(time.Hour)
	global.DB = db
}
