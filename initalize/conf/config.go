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
		Debug bool        `json:"debug"            desc:"是否开启Debug模式"`
		Mysql mysqlConfig `json:"mysql"            desc:"mysql配置"`
		Redis redisConfig `json:"redis"            desc:"redis"`
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
