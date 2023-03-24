package views

import (
	"book/initalize/database/mysql/book"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetAllBookData
// @Router /book/getAllData[get]
func GetAllBookData(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
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

func DelBookData(c *gin.Context) {
	isbn := c.Params.ByName("isbn")
	err := book.DelBook(isbn)
	if err != nil {
		response.Error(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)
}
