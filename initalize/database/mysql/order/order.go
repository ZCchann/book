package order

import (
	"book/initalize/database/mysql"
	"fmt"
	"time"
)

//新建订单
func CreateOrder(uuid, name, telephone, address string) (number string, err error) {
	TimeStamp := time.Now().Unix()
	var order createOrder
	order.UUID = uuid
	order.CreateTime = TimeStamp
	order.Addressee = name
	order.Address = address
	order.Telephone = telephone

	err = mysql.Mysql().Table("orderform").Create(&order).Error
	if err != nil {
		err = fmt.Errorf("orderform插入数据错误 请检查: %s", err)
		return
	}
	number, err = getOrderNumber(TimeStamp, uuid)
	if err != nil {
		return "", err
	}
	return number, nil

}

// getOrderNumber 获取指定时间的订单编号
func getOrderNumber(createTime int64, uuid string) (number string, err error) {
	err = mysql.Mysql().Table("orderform").Where("create_time=? AND uuid=?", createTime, uuid).First(&number).Error
	if err != nil {
		err = fmt.Errorf("getOrderNumber 查询orderform错误 请检查: %s", err)
		return
	}
	return

}

// AddOrderList 新增订单详情
func AddOrderList(orderList []OrderList, orderNumber string) (err error) {
	timestamp := time.Now().Unix()
	for _, i := range orderList {
		i.Number = orderNumber
		i.CreateTime = timestamp
	}
	err = mysql.Mysql().Table("orderlist").Create(&orderList).Error
	if err != nil {
		err = fmt.Errorf("orderlist插入数据错误 请检查: %s", err)
		return
	}
	return nil
}

// GetOrderList 获取订单信息
func GetOrderList(uuid string) (result []OrderForm, err error) {
	err = mysql.Mysql().Table("orderform").Where("uuid=?", uuid).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetOrderList 查询orderform错误 请检查: %s", err)
		return
	}
	return
}

func GetOrderDetails(orderNumber string) (result []OrderDetails, err error) {
	err = mysql.Mysql().Joins("JOIN bookdata ON orderlist.bookid = bookdata.id").Table("orderlist").
		Select("orderlist.number, orderlist.amount, orderlist.total_price, bookdata.title, bookdata.type, bookdata.price, bookdata.isbn, bookdata.publicationDate, bookdata.press").
		Where("orderlist.number = ?", orderNumber).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetOrderDetails 查询orderlist错误 请检查: %s", err)
		return
	}
	return

}

// GetAllOrderList 获取所有订单信息
func GetAllOrderList() (result []OrderForm, err error) {
	err = mysql.Mysql().Table("orderform").Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetAllOrderList 查询orderform错误 请检查: %s", err)
		return
	}
	return
}

// ExportOrderData 导出订单信息
func ExportOrderData(orderListID string) (result []ExportBookData, err error) {
	err = mysql.Mysql().Table("orderlist").
		Select("SUM(orderlist.amount) as total_amount, bookdata.title, bookdata.isbn, bookdata.press, bookdata.type, bookdata.restriction").
		Joins("INNER JOIN bookdata ON orderlist.bookid = bookdata.id").
		Where("orderlist.number IN (?)", orderListID).
		Group("orderlist.bookid, bookdata.id").
		Scan(&result).Debug().Error
	if err != nil {
		err = fmt.Errorf("ExportOrderData 导出订单信息错误 请检查: %s", err)
		return nil, err
	}
	return
}
