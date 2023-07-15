package main

import (
	"book/api/router"
	"book/initalize/conf"
	"book/initalize/database"
	"book/initalize/message"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if os.Getenv("Debug") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	file := "./conf/config.json"
	log.Println("加载配置文件:", file)
	conf.Init(file)

	log.Println("初始化数据库")
	database.InitMysql()
	database.InitRedis()

	log.Println("初始化路由")
	r := gin.Default()
	router.Register(r)

	// 判断一下LINE机器人是否启用
	if conf.Conf().LineBot.State {
		log.Println("初始化Line机器人")
		err := message.Line().InitLine()
		if err != nil {
			log.Println("初始化Line机器人错误 请检查: ", err)
			return
		}
	}

	log.Println("启动监听")
	err := r.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}

}
