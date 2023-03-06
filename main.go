package main

import (
	"book/api/router"
	"book/db/init"
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

	log.Println("初始化路由")
	r := gin.Default()
	router.Register(r)
	init.InitDB()

	log.Println("启动监听")
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
