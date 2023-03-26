package view

import (
	"book/initalize/database/mysql/user"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUser(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	if request.Username == "" {
		response.BadRequest(c, "`username` is required")
		return
	}
	if request.Password == "" {
		response.BadRequest(c, "`password` is required")
		return
	}

	if err := user.AddUser(request.Username, request.Password); err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

func UpdatePassword(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	if request.Username == "" {
		response.BadRequest(c, "`username` is required")
		return
	}
	if request.Password == "" {
		response.BadRequest(c, "`password` is required")
		return
	}

	err := user.UpdateUserPassword(request.Username, request.Password)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

func DeleteUser(c *gin.Context) {
	username := c.Params.ByName("username")
	if username == "" {
		response.BadRequest(c, "`username` is required")
		return
	}

	err := user.DelUser(username)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

func GetAllUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := user.GetAllUser()
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	sPage := page*pageSize - pageSize //起始index
	ePage := page * pageSize          //结束index
	total := len(res)                 //数据总页数

	//分页
	if len(res) < pageSize-1 {
		response.DataWtihPage(c, res, total)
	} else if ePage > total {
		ret := res[sPage:total]
		response.DataWtihPage(c, ret, total)
	} else {
		ret := res[sPage : ePage-1]
		response.DataWtihPage(c, ret, total)
	}
}
