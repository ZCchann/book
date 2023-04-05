package view

import (
	"book/initalize/database/mysql/user"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//AddUser 添加用户
func AddUser(c *gin.Context) {
	var request user.User
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := user.AddUser(request.Username, request.Password, request.Email)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)
}

func DeleteUser(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	if uuid == "" {
		response.BadRequest(c, "`uuid` is required")
		return
	}

	err := user.DelUser(uuid)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}

// @Route /user/getAllUser [GET]
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

// GetUser 获取用户信息
// @Route /getUser/:uuid/ [GET]
func GetUser(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	res, err := user.GetUserForID(uuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		log.Println(err.Error())
		return
	}
	response.Data(c, res)
}

//SearchUserData 搜索用户数据
func SearchUserData(c *gin.Context) {
	username := c.Query("username")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := user.SearchUser(username)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
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

//UpdateUser 更新用户信息
// @Route /updateUser [POST]
func UpdateUser(c *gin.Context) {
	var request user.User
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}

	if request.Password != "" {
		err := user.UpdateUserPassword(request.UUID, request.Email, request.Password)
		if err != nil {
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	} else {
		err := user.UpdateUserEmail(request.UUID, request.Email)
		if err != nil {
			log.Println(err)
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	}

}
