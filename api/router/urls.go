package router

import (
	"book/api/user"
	"book/api/views"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	user.RegisterRouter(r)

	group := r.Group("/book")
	{
		group.GET("/getAllData", views.GetAllBookData)
	}
}
