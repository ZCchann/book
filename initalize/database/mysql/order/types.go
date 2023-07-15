package order

type OrderForm struct {
	Number     string `json:"number"`
	CreateTime int    `json:"create_time"`
	Addressee  string `json:"addressee"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	AddressID  int    `json:"address_id"`
}

type OrderList struct {
	Number     string `json:"number"` // 订单编号
	BookID     int    `json:"id" gorm:"column:bookid"`
	Amount     int    `json:"amount"`
	TotalPrice int    `json:"total_price"`
	CreateTime int64  `json:"create_time"`
}

type SubmitOrder struct {
	OrderData []OrderList `json:"order_data"`
	Addressee string      `json:"addressee"`
	Telephone string      `json:"telephone"`
	Address   string      `json:"address"`
}

type OrderDetails struct {
	Number          string `json:"number"`
	Amount          int    `json:"amount"`
	TotalPrice      int    `json:"total_price"`
	ISBN            string `json:"isbn"`                                           // 书ISBN号
	Title           string `json:"title"`                                          // 书名
	Type            string `json:"type"`                                           // 类型 漫画/小说
	Price           int    `json:"price"`                                          // 定价
	PublicationDate string `json:"publication_date" gorm:"column:publicationDate"` // 出版日
	Press           string `json:"press"`                                          // 出版社
}

type ExportBookData struct {
	Title       string `json:"title"`        // 书名
	ISBN        string `json:"isbn"`         // 书ISBN号
	Press       string `json:"press"`        // 出版社
	Type        string `json:"type"`         // 类型 漫画/小说
	Restriction int    `json:"restriction"`  // 判断是否为限制级 1为是限制级
	TotalAmount int    `json:"total_amount"` // 合并后的总数
}

// 新建订单使用
type createOrder struct {
	UUID       string `json:"uuid"`
	Addressee  string `json:"addressee"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	CreateTime int64  `json:"create_time"`
}
