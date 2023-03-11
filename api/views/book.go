package views

import (
	"book/initalize/database/mysql/book"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetAllBookData
// @Router /attack/getalldata/ [get]
func GetAllBookData(c *gin.Context) {
	res, err := book.GetAllBook()
	if err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	response.Data(c, res)

}
