package router

import (
	"book/api/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	user.RegisterRouter(r)

}
