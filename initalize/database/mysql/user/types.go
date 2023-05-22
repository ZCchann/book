package user

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	UUID        string `json:"uuid"`
	AuthorityID int    `json:"authorityid"`
}

type UpdateUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	NewEmail string `json:"new_email"`
	UUID     string `json:"uuid"`
	Code     string `json:"code"`
}
