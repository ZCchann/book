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

	log.Println("初始化路由")
	r := gin.Default()
	router.Register(r)

	log.Println("初始化Line机器人")
	err := message.Line().InitLine()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("启动监听")
	err = r.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}

}
