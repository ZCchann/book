package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"sync"
)

type (
	conf struct {
		Debug      bool        `json:"debug"            desc:"是否开启Debug模式"`
		Mysql      mysqlConfig `json:"mysql"            desc:"mysql配置"`
		Redis      redisConfig `json:"redis"            desc:"redis"`
		LineBot    lineBot     `json:"line_bot"         desc:"line机器人相关"`
		MailServer mail        `json:"mail_server"             desc:"邮件相关"`
	}
	mysqlConfig struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
	}
	redisConfig struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Database int    `json:"database"`
	}
	lineBot struct {
		ChannelSecret      string `json:"channel_secret"       desc:"频道secret token" `
		ChannelAccessToken string `json:"channel_access_token" desc:"access token"`
		GroupID            string `json:"group_id"`
	}
	mail struct {
		Mail       string `json:"mail"         desc:"邮箱地址"`
		Password   string `json:"password"     desc:"邮箱密码/授权码"`
		SmtpServer string `json:"smtp_server"  desc:"SMTP服务器地址"`
		SmtpPort   int    `json:"smtp_port"    desc:"SMTP端口号"`
	}
)

var (
	c    = new(conf)
	lock = new(sync.RWMutex)
)

func Init(file string) {
	err := BindJSON(file, &c)
	if err != nil {
		log.Fatalln("初始化配置文件失败：", err)
	}
}

func Conf() *conf {
	lock.RLock()
	defer lock.RUnlock()
	return c
}

func BindJSON(file string, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("`v` must be a pointer")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	return err
}
