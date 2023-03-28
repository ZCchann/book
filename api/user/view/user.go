package view

import (
	"book/initalize/database/mysql/user"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// AddUser
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

// UpdatePassword 更新用户密码
//func UpdatePassword(c *gin.Context) {
//	var request loginRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		response.Error(c, "ShouldBindJSON："+err.Error())
//		return
//	}
//	if request.Username == "" {
//		response.BadRequest(c, "`username` is required")
//		return
//	}
//	if request.Password == "" {
//		response.BadRequest(c, "`password` is required")
//		return
//	}
//
//	err := user.UpdateUserPassword(request.Username, request.Password)
//	if err != nil {
//		response.Error(c, err.Error())
//		return
//	}
//	response.Success(c)
//}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		response.BadRequest(c, "`id` is required")
		return
	}

	err := user.DelUser(id)
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

func GetOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	res, err := user.GetUserForID(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, res)
}

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

//UpdateUser
// Route /updateUser [POST]
func UpdateUser(c *gin.Context) {
	var request user.User
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	log.Println(request.Id)
	log.Println(request.Email)
	log.Println(request.Password)
	if request.Password != "" {
		log.Println("1")
		err := user.UpdateUserPassword(request.Id, request.Email, request.Password)
		if err != nil {
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	} else {
		log.Println("2")
		err := user.UpdateUserEmail(request.Id, request.Email)
		if err != nil {
			log.Println(err)
			response.Error(c, err.Error())
			return
		}
		response.Success(c)
	}

}
