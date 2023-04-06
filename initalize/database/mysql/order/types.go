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
	BookID     int `json:"id"`
	Amount     int `json:"amount"`
	TotalPrice int `json:"total_price"`
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
	ISBN            string `json:"isbn"`             // 书ISBN号
	Title           string `json:"title"`            // 书名
	Type            string `json:"type"`             // 类型 漫画/小说
	Price           int    `json:"price"`            // 定价
	PublicationDate string `json:"publication_date"` // 出版日
	Press           string `json:"press"`            // 出版社
}
