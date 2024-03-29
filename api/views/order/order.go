package order

import (
	"book/initalize/database/mysql/order"
	"book/pkg/message/line"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// getPageData 返回分页数据
func getPageData(data []order.OrderForm, page int, pageSize int) []order.OrderForm {
	start := (page - 1) * pageSize
	end := page * pageSize
	if start > len(data) {
		return []order.OrderForm{}
	}
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}

//CreateOrder 新建订单
// @Route /order/createOrder
func CreateOrder(c *gin.Context) {
	var request order.SubmitOrder
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, "ShouldBindJSON："+err.Error())
		return
	}
	uuid := c.GetHeader("uuid")

	orderNumber, err := order.CreateOrder(uuid, request.Addressee, request.Telephone, request.Address)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}

	errList := order.AddOrderList(request.OrderData, orderNumber)
	if errList != nil {
		response.Error(c, errList.Error())
		return
	}
	response.Success(c)

	line.SendMessage(fmt.Sprintf("用户提交一笔新订单 请注意查看 订单ID: %s", orderNumber))

}

//GetOrder 获取uuid 所有的订单信息
// @Route /order/get_order [GET]
func GetOrder(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := order.GetOrderList(uuid)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	total := len(res) //数据总页数
	ret := getPageData(res, page, pageSize)
	response.DataWtihPage(c, ret, total)
}

// GetAllOrder 获取uuid 所有的订单信息
// @Route /order/get_all_order [GET]
func GetAllOrder(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	res, err := order.GetAllOrderList()
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	total := len(res) //数据总页数
	ret := getPageData(res, page, pageSize)
	response.DataWtihPage(c, ret, total)

}

func GetOrderDetails(c *gin.Context) {
	orderNumber := c.Query("number")
	requests, err := order.GetOrderDetails(orderNumber)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, requests)
}

// ExportOrderData 返回订单详情
func ExportOrderData(c *gin.Context) {
	orderNumberList := c.Query("order_number_list")
	res, err := order.ExportOrderData(orderNumberList)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, res)
}
