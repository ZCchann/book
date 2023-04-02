package view

type loginRequest struct {
	//Addressee     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	//Remark   string `json:"remark" form:"remark"`
}
