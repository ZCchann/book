package permissions

import (
	"book/initalize/database/mysql/authority"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func GetRoute(c *gin.Context) {
	var result []Routers
	result = append(result, AdminMenu())
	result = append(result, OrderMenu())
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
	log.Println(result)
	response.Data(c, result)
}

func GetPermissionsByID(c *gin.Context) {
	permissionsID := c.Query("permissions_id")
	result, err := authority.GetPermissionsByID(permissionsID)
	if err != nil {
		log.Println(err)
		return
	}

	d := new(authority.EditPermissions)
	d.ID = result.ID
	d.RuleName = result.RuleName

	// 遍历一下mysql返回值 重组数据
	v := reflect.ValueOf(result)
	for i := 0; i < v.NumField(); i++ {
		status := v.Field(i).Interface()
		name := v.Type().Field(i).Name
		if name == "ID" || name == "RuleName" {
			// 跳过这两个ID名称
			continue
		}
		s := false //临时变量 用于接收遍历结构体的布尔值
		var value authority.Permission

		value.Name = name
		if boolValue, ok := status.(bool); ok {
			s = boolValue
		}
		value.State = s
		d.Permissions = append(d.Permissions, value)

	}

	response.Data(c, d)
}

func UpdatePermissionsByID(c *gin.Context) {
	var request authority.EditPermissions
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := authority.UpdatePermissionsByID(request)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}
