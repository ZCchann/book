package views

import (
	"book/initalize/database/mysql/book"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
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

func DelBookData(c *gin.Context) {
	isbn := c.Params.ByName("isbn")
	err := book.DelBook(isbn)
	if err != nil {
		response.Error(c, err.Error())
		log.Fatal(err.Error())
		return
	}
	response.Success(c)
}
