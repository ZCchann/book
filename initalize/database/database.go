package database

import (
	"book/initalize/conf"
	"book/initalize/database/mysql"
	"book/initalize/database/redis"
	"log"
)

func InitMysql() {
	mysql.Mysql().InitDB(
		conf.Conf().Mysql.Username,
		conf.Conf().Mysql.Password,
		conf.Conf().Mysql.Host,
		conf.Conf().Mysql.Port,
		conf.Conf().Mysql.Database,
	)
}

func InitRedis() {
	err := redis.Redis().Connect(
		conf.Conf().Redis.Addr,
		conf.Conf().Redis.Password,
		conf.Conf().Redis.Database,
	)
	if err != nil {
		log.Fatalln("连接Redis失败:", err)
	}
}
