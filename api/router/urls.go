package router

import (
	"book/api/user"
	"book/api/views/bookData"
	"book/api/views/order"
	"book/pkg/grf/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(middleware.Cross()) //解决前端跨域请求报错问题
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	user.RegisterRouter(r)

	book := r.Group("/book")
	{
		book.GET("/getAllData", bookData.GetAllBookData)
		book.GET("/getData/:id/", bookData.GetBookData)
		book.GET("/search/", bookData.SearchBookData)
		book.POST("/addData", bookData.AddBookData)
		book.POST("/editData", bookData.EditBookData)
		book.POST("/fileUpdate", bookData.FileUpdate)
		book.DELETE("/delData/:id/", bookData.DelBookData)
	}

	orderList := r.Group("/order")
	{
		orderList.GET("/get_order", order.GetOrder)
		orderList.GET("/get_all_order", order.GetAllOrder)
		orderList.GET("/get_order_details", order.GetOrderDetails)
		orderList.POST("/create", order.CreateOrder)

	}

}
