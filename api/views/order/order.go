package order

import (
	"book/initalize/database/mysql/order"
	"book/pkg/message/line"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

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
		errNumber := len(errList)
		var errStr string
		for _, i := range errList {
			errStr += i
		}
		response.Error(c, fmt.Sprintf("失败%d条数据,失败信息: %d", errNumber, errStr))
		return
	}
	response.Success(c)

	line.SendMessage(fmt.Sprintf("用户提交一笔新订单 请注意查看 订单ID: %s", orderNumber))

}

//GetOrder 获取uuid 所有的订单信息
// @Route /order/get_order [GET]
func GetOrder(c *gin.Context) {
	uuid := c.GetHeader("uuid")
	requests, err := order.GetOrderList(uuid)
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, requests)
}

// GetAllOrder 获取uuid 所有的订单信息
// @Route /order/get_all_order [GET]
func GetAllOrder(c *gin.Context) {
	requests, err := order.GetAllOrderList()
	if err != nil {
		response.BadRequest(c, err.Error())
		log.Println(err.Error())
		return
	}
	response.Data(c, requests)
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
