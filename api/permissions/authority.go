package permissions

import (
	"book/initalize/database/mysql/authority"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func GetRoute(c *gin.Context) {
	var result []Routers
	result = append(result, AdminMenu())
	result = append(result, OrderMenu())
	response.Data(c, result)

}

func GetPermissions(c *gin.Context) {
	result, err := authority.GetPermissionsGroup()
	if err != nil {
		log.Println(err)
		return
	}
	for _, i := range result {
		fmt.Println(i.Name)

	}
}

func GetAllPermissionsIDName(c *gin.Context) {
	result, err := authority.GetAllPermissionsIDName()
	if err != nil {
		log.Println(err)
		return
	}
	response.Data(c, result)
}

func GetPermissionsByID(c *gin.Context) {
	permissionsID := c.Query("permissions_id")
	result, err := authority.GetPermissionsByID(permissionsID)
	if err != nil {
		log.Println(err)
		return
	}
	response.Data(c, result)
}
