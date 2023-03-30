package views

import (
	"book/initalize/database/mysql/book"
	"book/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
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

// DelBookData @Router /book/delData/[delete]
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

// GetBookData
// @Router /book/getData/:id/[GET]
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

// EditBookData
// @Router /book/editData[POST]
func EditBookData(c *gin.Context) {
	var request book.BookData
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	log.Println(request)
	err := book.EditBook(request)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)

}

// AddBookData
// @Router /book/addData[POST]
func AddBookData(c *gin.Context) {
	var request book.BookData

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	log.Println(request)
	err := book.AddBook(request)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Success(c)

}

// SearchBookData
// @Router /book/search/?title=bookTitle [GET]
func SearchBookData(c *gin.Context) {
	title := c.Query("title")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := book.SearchBook(title)
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

func FileUpdate(c *gin.Context) {
	f, _ := c.FormFile("file")
	log.Println(f.Filename)
	//SaveUploadedFile上传表单文件到指定的路径
	err := c.SaveUploadedFile(f, "./file/"+f.Filename)
	if err != nil {
		log.Println("err ", err)
		return
	}
	result, err := ParsingFile("./file/" + f.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	response.Data(c, result)
}

func ParsingFile(filePath string) (result []book.BookData, err error) {
	var fileData book.BookData
	var retData []book.BookData
	file, err := xlsx.FileToSlice(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sheet := file[0] //只取文件的sheet0
	//第一行(0)为标题 标题不导入 从1开始
	for i := 1; i < len(sheet); i++ {
		fileData.ISBN = sheet[i][0]
		fileData.Tittle = sheet[i][1]
		fileData.Price, _ = strconv.Atoi(sheet[i][2])
		fileData.Press = sheet[i][3]
		fileData.Type = sheet[i][4]
		fileData.Restriction, _ = strconv.Atoi(sheet[i][5])
		fileData.Author = sheet[i][6]
		fileData.PublicationDate = sheet[i][7]
		retData = append(retData, fileData)
	}
	return retData, nil
}
