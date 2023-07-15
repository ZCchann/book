package book

import (
	"book/initalize/database/mysql"
	"fmt"
)

type BookData struct {
	ID              int    `json:"id"`
	ISBN            string `json:"isbn"`                                           // 书ISBN号
	Title           string `json:"title"`                                          // 书名
	Price           int    `json:"price"`                                          // 定价
	Press           string `json:"press"`                                          // 出版社
	Type            string `json:"type"`                                           // 类型 漫画/小说
	Restriction     int    `json:"restriction"`                                    // 判断是否为限制级 1为是限制级
	Author          string `json:"author"`                                         // 作者
	PublicationDate string `json:"publication_date" gorm:"column:publicationDate"` // 出版日
}

//GetBook 获取单条图书数据
func GetBook(ID string) (result BookData, err error) {
	err = mysql.Mysql().Table("bookdata").Where("id=?", ID).First(&result).Error
	if err != nil {
		err = fmt.Errorf("GetBook 查询bookdata错误 请检查: %s", err)
		return
	}
	return
}

func GetAllBook() (result []BookData, err error) {
	err = mysql.Mysql().Table("bookdata").Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetAllBook 查询bookdata错误 请检查: %s", err)
		return
	}
	return
}

//AddBook 新增图书数据
func AddBook(data BookData) (err error) {
	err = mysql.Mysql().Table("bookdata").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("bookdata插入数据错误 请检查: %s", err)
		return
	}
	return nil

}

func DelBook(ID int) (err error) {
	err = mysql.Mysql().Table("bookdata").Where("id=?", ID).Delete(&BookData{}).Error
	if err != nil {
		err = fmt.Errorf("bookdata表删除数据错误 请检查: %s", err)
		return
	}
	return nil

}

func EditBook(data BookData) (err error) {
	err = mysql.Mysql().Table("bookdata").Where("id=?", data.ID).Save(&data).Error
	if err != nil {
		err = fmt.Errorf("bookdata update错误 请检查: %s", err)
		return
	}
	return
}

func SearchBook(title, startTime, endTime string) (result []BookData, err error) {
	if startTime != "" && title != "" {
		err = mysql.Mysql().Table("bookdata").Where("publicationDate BETWEEN ? AND ?", startTime, endTime).Where("title REGEXP ?", title).Find(&result).Error
		if err != nil {
			err = fmt.Errorf("SearchBook 查询bookdata错误 请检查: %s", err)
			return
		}
	}
	if startTime == "" && title != "" {
		err = mysql.Mysql().Table("bookdata").Where("title REGEXP ?", title).Find(&result).Error
		if err != nil {
			err = fmt.Errorf("SearchBook 查询bookdata错误 请检查: %s", err)
			return
		}
	}
	if startTime != "" && title == "" {
		err = mysql.Mysql().Table("bookdata").Where("publicationDate BETWEEN ? AND ?", startTime, endTime).Find(&result).Error
		if err != nil {
			err = fmt.Errorf("SearchBook 查询bookdata错误 请检查: %s", err)
			return
		}
	}

	return result, nil
}
