package mysql

import (
	"book/initalize/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var m *gorm.DB

func Mysql() *gorm.DB {
	return m
}

func InitDB() {
	username := conf.Conf().Mysql.Username
	password := conf.Conf().Mysql.Password
	host := conf.Conf().Mysql.Host
	port := conf.Conf().Mysql.Port
	database := conf.Conf().Mysql.Database
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(sqlUrl), nil)
	if err != nil {
		log.Fatal("连接mysql数据库错误 请检查: ", err)
		return
	}
	m = db
}
