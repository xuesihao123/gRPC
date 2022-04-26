package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMysql()(err error){

	dsn := "root:123456@(127.0.0.1:3307)/restaurant?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//Db , err = gorm.Open("mysql", "root:123456@(127.0.0.1:3307)/restaurant?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		return
	}

	return nil
}
