package util

import (
	util2 "book/pkg/grf/util"
	"book/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// SendVerificationCode 发送验证码
func SendVerificationCode(c *gin.Context) {
	email := c.Query("email")
	fmt.Println(email)
	err := util2.SendMail(email, "系统测试", "您正在修改邮箱 本次验证码为: "+util2.CreateCode(email))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c)
}
