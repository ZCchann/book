package order

import (
	"book/initalize/database/mysql"
	"fmt"
	"log"
	"strconv"
	"time"
)

//新建订单
func CreateOrder(uuid, name, telephone, address string) (number string, err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO orderform (`uuid`,`create_time`,`addressee`,`telephone`,`address`) VALUE (?,?,?,?,?)")
	if err != nil {
		log.Println("prepare fail")
		return "", err
	}

	// 获取时间戳
	TimeStamp := time.Now().Unix()

	// 传参到sql中执行
	_, err = stmt.Exec(uuid, TimeStamp, name, telephone, address)
	if err != nil {
		log.Println("exec fail")
		return "", err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return "", err
	}
	number, err = getOrderNumber(TimeStamp, uuid)
	if err != nil {
		log.Println("get number error: ", err)
		return "", err
	}
	return number, nil

}

// getOrderNumber 获取指定时间的订单编号
func getOrderNumber(createTime int64, uuid string) (number string, err error) {
	err = mysql.Mysql().DB.QueryRow("SELECT number FROM orderform WHERE create_time = ? and uuid=? ", createTime, uuid).Scan(&number)
	if err != nil {
		log.Println(err)
		return
	}
	return number, err
}

// AddOrderList 新增订单详情
func AddOrderList(orderList []OrderList, orderNumber string) (err []string) {
	var errList []string

	for i := 0; i < len(orderList); i++ {
		tx, err := mysql.Mysql().DB.Begin()
		if err != nil {
			log.Println("tx fail")
		}
		bookID := orderList[i].BookID
		amount := orderList[i].Amount
		totalPrice := orderList[i].TotalPrice
		timestamp := time.Now().Unix()
		// 准备sql语句
		stmt, err := tx.Prepare("INSERT INTO orderlist (`number`,`bookid`,`amount`,`total_price`,`create_time`) VALUE (?,?,?,?,?)")
		if err != nil {
			log.Println("prepare fail")
			errList = append(errList, fmt.Sprintf("bookID %s error: %s"), strconv.Itoa(bookID), err.Error())
			continue
		}
		// 传参到sql中执行
		_, err = stmt.Exec(orderNumber, bookID, amount, totalPrice, timestamp)
		if err != nil {
			log.Println("exec fail")
			errList = append(errList, fmt.Sprintf("bookID %s error: %s"), strconv.Itoa(bookID), err.Error())
			continue
		}
		// 提交
		err = tx.Commit()
		if err != nil {
			log.Println("commit error ", err)
			errList = append(errList, fmt.Sprintf("bookID %s error: %s"), strconv.Itoa(bookID), err.Error())
			continue
		}
	}

	if len(errList) > 0 {
		return errList
	}
	return nil
}

// GetOrderList 获取订单信息
func GetOrderList(uuid string) (result []OrderForm, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT number,addressee,telephone,address,create_time from orderform where  uuid = ?;", uuid)
	if err != nil {
		log.Println("1 ", err)
		return
	}

	for rows.Next() {
		var f OrderForm
		err = rows.Scan(&f.Number, &f.Addressee, &f.Telephone, &f.Address, &f.CreateTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}
	return result, err
}

func GetOrderDetails(orderNumber string) (result []OrderDetails, err error) {
	rows, err := mysql.Mysql().DB.Query("select orderlist.number,orderlist.amount,orderlist.total_price,bookdata.title,bookdata.type,bookdata.price, bookdata.isbn,bookdata.publicationDate,bookdata.press from orderlist, bookdata where orderlist.bookid=bookdata.id and number=?;", orderNumber)
	if err != nil {
		log.Println("1 ", err)
		return
	}

	for rows.Next() {
		var f OrderDetails
		err = rows.Scan(&f.Number, &f.Amount, &f.TotalPrice, &f.Title, &f.Type, &f.Price, &f.ISBN, &f.PublicationDate, &f.Press)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}
	return result, err
}

// GetAllOrderList 获取所有订单信息
func GetAllOrderList() (result []OrderForm, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT number,addressee,telephone,address,create_time from orderform;")
	if err != nil {
		log.Println("1 ", err)
		return
	}

	for rows.Next() {
		var f OrderForm
		err = rows.Scan(&f.Number, &f.Addressee, &f.Telephone, &f.Address, &f.CreateTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}
	return result, err
}
