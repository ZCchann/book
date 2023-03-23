package router

import (
	"book/api/user"
	"book/api/views"
	"book/pkg/grf/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(middleware.Cross()) //解决前端跨域请求报错问题
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	user.RegisterRouter(r)

	group := r.Group("/book")
	{
		group.GET("/getAllData", views.GetAllBookData)
		group.DELETE("/delData/:isbn/", views.DelBookData)
	}
}
