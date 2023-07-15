package util

import (
	"book/initalize/conf"
	"fmt"
	"gopkg.in/gomail.v2"
)

// SendMail 发送邮件
func SendMail(toMailAddress string, Header string, Body string) (err error) {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", conf.Conf().MailServer.Mail)
	//接收人
	m.SetHeader("To", toMailAddress)
	//主题
	m.SetHeader("Subject", Header)
	//内容
	m.SetBody("text/html", Body)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(conf.Conf().MailServer.SmtpServer, conf.Conf().MailServer.SmtpPort, conf.Conf().MailServer.Mail, conf.Conf().MailServer.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		return err
	}

	return nil
}
