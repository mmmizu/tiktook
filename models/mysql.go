package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB

type database struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}

func DbInit() {
	myDB := &database{
		username: "tiktok",
		password: "2aebEzEnWFEi44k2",
		host:     "47.113.178.29",
		port:     "3300",
		dbName:   "demo",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		myDB.username, myDB.password, myDB.host, myDB.port, myDB.dbName)
	tempDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return
	}
	sqlDB, err := tempDb.DB()
	if err != nil {
		log.Printf("database setup error %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          //最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //设置连接空闲超时
	db = tempDb
}

func GetDb() *gorm.DB {
	return db
}
