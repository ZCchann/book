package user

import (
	"book/api/user/view"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", view.Login)

	group := r.Group("/user")
	{
		group.GET("/getAllUser", view.GetAllUser)
		group.GET("/getUser/:uuid/", view.GetOneUser)
		group.GET("/search/", view.SearchUserData)
		group.POST("/addUser", view.AddUser)
		group.POST("/updateUser", view.UpdateUser)
		group.DELETE("/delUser/:uuid/", view.DeleteUser)
	}
}
