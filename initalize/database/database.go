package database

import (
	"book/initalize/conf"
	"book/initalize/database/mysql"
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
