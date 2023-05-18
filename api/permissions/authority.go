package permissions

import (
	"book/initalize/database/mysql/authority"
	"book/initalize/database/mysql/user"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

// GetRoute 返回动态路由给前端
func GetRoute(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	res, err := PermissionFiltering(uuid)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Data(c, res)

}

// GetPermissionDemo 返回一个空的权限组框架 用于前端添加新的权限信息使用
func GetPermissionDemo(c *gin.Context) {
	result, err := authority.GetPermissionsFirst()
	if err != nil {
		log.Println(err)
		return
	}
	d := new(authority.EditPermissions)
	d.ID = 0
	d.RuleName = ""

	// 遍历一下mysql返回值 重组数据
	v := reflect.ValueOf(result)
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		if name == "ID" || name == "RuleName" {
			// 跳过这两个ID名称
			continue
		}
		var value authority.Permission
		value.Name = name
		value.State = false
		d.Permissions = append(d.Permissions, value)
	}
	response.Data(c, d)
}

// GetAllPermissionsIDName 返回权限ID与权限组名称、权限详情信息
func GetAllPermissionsIDName(c *gin.Context) {
	result, err := authority.GetAllPermissionsIDName()
	if err != nil {
		log.Println(err)
		response.Error(c, err.Error())
		return
	}
	response.Data(c, result)
}

func GetPermissionsByID(c *gin.Context) {
	permissionsID := c.Query("permissions_id")
	result, err := authority.GetPermissionsByID(permissionsID)
	if err != nil {
		log.Println(err)
		response.Error(c, err.Error())
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

// AddPermission 添加权限
func AddPermission(c *gin.Context) {
	var request authority.EditPermissions
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := authority.AddPermission(request)
	if err != nil {
		log.Println(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

//UpdatePermissionsByID 通过权限组ID来更改权限内容
func UpdatePermissionsByID(c *gin.Context) {
	var request authority.EditPermissions
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := authority.UpdatePermissionsByID(request)
	if err != nil {
		log.Println(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

func DeletePermissionsByID(c *gin.Context) {
	var request []authority.EditPermissions
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("json eof", err)
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	log.Println(request[0].ID)
	//删除前检查 若有用户在使用将要删除的权限ID 将返回错误
	for i := 0; i < len(request); i++ {
		id := request[i].ID
		res, err := user.GetUserPermissionForID(id)
		if err != nil {
			log.Println(err)
			response.Error(c, "ShouldBindJSON："+err.Error())
			return
		}
		if len(res) != 0 {
			response.Error(c, fmt.Sprintf("已成功删除 %d 个权限组, 权限组ID %d 被占用,请解除占用后再提交删除", i, id))
			return
		}
		err = authority.DeletePermissionByID(id)
		if err != nil {
			log.Println(err)
			response.Error(c, "ShouldBindJSON："+err.Error())
			return
		}
	}
	response.Success(c)
}
