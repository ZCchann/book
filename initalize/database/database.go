package database

import (
	"book/initalize/conf"
	"book/initalize/database/mysql"
	"book/initalize/database/redis"
	"log"
)

func InitMysql() {
	mysql.InitDB()
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
