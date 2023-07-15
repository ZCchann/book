package util

import (
	"book/initalize/database/redis"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// CreateCode 生成验证码 存入redis
func CreateCode(mailAddress string) string {
	min := int32(100000) //设置随机数下限
	max := int32(999999) //设置随机数上限
	rand.Seed(time.Now().UnixNano())
	num := rand.Int31n(max-min) + min
	// 生成一个随机数做为验证码 有效期10分钟
	redis.Redis().Set(context.Background(), fmt.Sprintf("code:%s", mailAddress), num, 600*time.Second)
	return strconv.Itoa(int(num))
}

// CheckCode 检查验证码
func CheckCode(mailAddress, code string) (err error) {
	res := redis.Redis().Get(context.Background(), fmt.Sprintf("code:%s", mailAddress))
	if code == res.Val() {
		return nil
	} else {
		err = fmt.Errorf("验证码错误")
		return
	}
}
