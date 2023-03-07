package user

import (
	"book/api/user/view"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {

	group := r.Group("/user")
	{
		group.POST("/adduser", view.AddUser)
		group.POST("/updatepasswd", view.UpdatePassword)
		group.DELETE("/:username/", view.DeleteUser)
	}
}
