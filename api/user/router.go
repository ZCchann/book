package user

import (
	"book/api/authority"
	"book/api/user/view"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", view.Login)

	group := r.Group("/user")
	{
		group.GET("/getAllUser", view.GetAllUser)
		group.GET("/getUser/:uuid/", view.GetUser)
		group.GET("/search/", view.SearchUserData)
		group.GET("/get_user_all_address", view.GetUserAllAddress)
		group.GET("/get_user_address", view.GetUserAddress)
		group.GET("/get_route", authority.GetRoute)
		group.POST("/addUser", view.AddUser)
		group.POST("/updateUser", view.UpdateUser)
		group.POST("/add_user_address", view.AddUserAddress)
		group.POST("/update_address", view.UpdateUserAddress)
		group.DELETE("/delUser/:uuid/", view.DeleteUser)
		group.DELETE("/del_address", view.DeleteUserAddress)
	}
}
