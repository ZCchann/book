package view

import (
	"book/initalize/database/mysql/user"
	"book/pkg/grf/util"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// getPageData 返回分页数据
func getPageData(data []user.User, page int, pageSize int) []user.User {
	start := (page - 1) * pageSize
	end := page * pageSize
	if start > len(data) {
		return []user.User{}
	}
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}

//AddUser 添加用户
func AddUser(c *gin.Context) {
	var request user.User
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := user.AddUser(request)
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

// GetAllUser 返回所有用户信息
// @Route /user/getAllUser [GET]
func GetAllUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := user.GetAllUser()
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	//分页
	total := len(res) //数据总页数
	ret := getPageData(res, page, pageSize)
	response.DataWtihPage(c, ret, total)
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

	//分页
	total := len(res) //数据总页数
	ret := getPageData(res, page, pageSize)
	response.DataWtihPage(c, ret, total)
}

//AdminUpdateUser 更新用户信息
// @Route /admin/updateUser [POST]
func AdminUpdateUser(c *gin.Context) {
	var request user.User
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	if request.Password != "" {
		err := user.UpdateUserPassword(request.UUID, request.Email, request.Password, "admin", request.AuthorityID)
		if err != nil {
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	} else {
		err := user.UpdateUserEmail(request.UUID, request.Email, "admin", request.AuthorityID)
		if err != nil {
			log.Println(err)
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	}

}

// UserUpdate 用户更新数据
func UserUpdate(c *gin.Context) {
	var request user.UpdateUser
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}

	// 检查验证码
	err := util.CheckCode(request.Email, request.Code)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	if request.Password != "" {
		err := user.UpdateUserPassword(request.UUID, request.Email, request.Password, "user", 0)
		if err != nil {
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	} else {
		err := user.UpdateUserEmail(request.UUID, request.NewEmail, "user", 0)
		if err != nil {
			log.Println(err)
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	}
}
