package user

import (
	"book/api/permissions"
	"book/api/user/view"
	"book/pkg/grf/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", view.Login)
	r.Use(middleware.Jwt())
	r.POST("/logout", view.Logout)

	group := r.Group("/user")
	{
		group.GET("/getAllUser", view.GetAllUser)
		group.GET("/getUser/:uuid/", view.GetUser)
		group.GET("/search/", view.SearchUserData)
		group.GET("/get_user_all_address", view.GetUserAllAddress)
		group.GET("/get_user_address", view.GetUserAddress)
		group.POST("/addUser", view.AddUser)
		group.POST("/updateUser", view.UpdateUser)
		group.POST("/add_user_address", view.AddUserAddress)
		group.POST("/update_address", view.UpdateUserAddress)
		group.DELETE("/delUser/:uuid/", view.DeleteUser)
		group.DELETE("/del_address", view.DeleteUserAddress)
	}
	permission := r.Group("/permissions")
	{
		permission.GET("/get_route/:uuid/", permissions.GetRoute)
		permission.GET("/get_permissions_id_name", permissions.GetAllPermissionsIDName)
		permission.GET("/get_permissions_by_id", permissions.GetPermissionsByID)
		permission.GET("/get_permissions_demo", permissions.GetPermissionDemo)
		permission.POST("/add_permission", permissions.AddPermission)
		permission.POST("/update_permissions_by_id", permissions.UpdatePermissionsByID)

	}
}
