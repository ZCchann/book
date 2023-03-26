package book

import (
	"book/initalize/database/mysql"
	"log"
	"time"
)

type BookData struct {
	ID              int       `json:"id"`
	ISBN            string    `json:"isbn"`             // 书ISBN号
	Tittle          string    `json:"tittle"`           // 书名
	Price           int       `json:"price"`            // 定价
	Press           string    `json:"press"`            // 出版社
	Type            string    `json:"type"`             // 类型 漫画/小说
	Restriction     int       `json:"restriction"`      // 判断是否为限制级 1为是限制级
	Author          string    `json:"author"`           // 作者
	PublicationDate time.Time `json:"publication_date"` // 出版日
}

//GetBook 获取单条图书数据
func GetBook(ID string) (result []BookData, err error) {
	var f BookData
	err = mysql.Mysql().DB.QueryRow("SELECT * FROM bookdata WHERE id = ? ", ID).Scan(&f.ID, &f.ISBN, &f.Tittle, &f.Price, &f.Press, &f.Type, &f.Restriction, &f.Author, &f.PublicationDate)
	if err != nil {
		log.Println(err)
		return
	}
	result = append(result, f)
	return result, err
}

func GetAllBook() (result []BookData, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT * FROM bookdata")
	if err != nil {
		log.Println(err)
		return result, err
	}

	for rows.Next() {
		var f BookData
		err = rows.Scan(&f.ID, &f.ISBN, &f.Tittle, &f.Price, &f.Press, &f.Type, &f.Restriction, &f.Author, &f.PublicationDate)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err

}

//AddBook 新增图书数据
func AddBook(data BookData) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}
	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO bookdata (`isbn`,`title`,`price`,`press`,`Type`,`restriction`,`author`,`publicationDate`) VALUE (?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(data.ISBN, data.Tittle, data.Price, data.Press, data.Type, data.Restriction, data.Author, data.PublicationDate)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return err
	}
	return nil
}

func DelBook(ID string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM bookdata WHERE id = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
	}
	return nil
}

func EditBook(data BookData) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return err
	}
	stmt, err := tx.Prepare("UPDATE bookdata SET isbn=? ,title=? ,price=? ,press=? ,type=? ,restriction=? ,author=? WHERE id=?")
	if err != nil {
		log.Println("Prepare fail ", err)
		return err
	}
	_, err = stmt.Exec(data.ISBN, data.Tittle, data.Price, data.Press, data.Type, data.Restriction, data.Author, data.ID)
	if err != nil {
		log.Println("exec fail ", err)
		return err
	}
	tx.Commit()
	return nil
}
