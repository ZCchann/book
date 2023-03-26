package views

import (
	"book/initalize/database/mysql/book"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetAllBookData 按照前端所需数量 返回数据库中数据给前端
// @Router /book/getAllData[get]
func GetAllBookData(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := book.GetAllBook()
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

// @Royter /book/delData/:id/[delete]
func DelBookData(c *gin.Context) {
	id := c.Params.ByName("id")
	err := book.DelBook(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)
}

func GetBookData(c *gin.Context) {
	id := c.Params.ByName("id")
	res, err := book.GetBook(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, res)
}

func EditBookData(c *gin.Context) {
	var request book.BookData
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}

	err := book.EditBook(request)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)

}

func AddBookData(c *gin.Context) {
	var request book.BookData
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	err := book.AddBook(request)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)
}
