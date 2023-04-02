package view

import (
	"book/initalize/database/mysql/user"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
)

// AddUserAddress 添加用户地址
func AddUserAddress(c *gin.Context) {
	var request user.Address
	uuid := c.GetHeader("uuid")
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := user.AddUserAddress(request, uuid)
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}

	response.Success(c)

}

func GetUserAllAddress(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	result, err := user.GetUserAllAddress(uuid)
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	response.Data(c, result)
}

func GetUserAddress(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	AddressId := c.Query("address_id")
	result, err := user.GetUserAddress(uuid, AddressId)
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	response.Data(c, result)
}

// UpdateUserAddress 更新用户地址信息
// @Route /user/update_address [POST]
func UpdateUserAddress(c *gin.Context) {
	var request user.Address
	uuid := c.GetHeader("uuid")
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := user.UpdateUserAddress(request, uuid)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)

}

func DeleteUserAddress(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	AddressId := c.Query("address_id")
	log.Println("uuid ", uuid, " addressID ", AddressId)
	err := user.DeleteUserAddress(AddressId, uuid)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)

}
