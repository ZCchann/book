package token

import (
	"book/initalize/database/redis"
	"book/pkg/grf/jwt"
	"context"
	"errors"
	"time"
)

func CreateToken(username string) (string, error) {
	var err error

	atClaims := jwt.CustomClaims{}
	atClaims.Username = username
	atClaims.ExpiresAt = time.Now().Add(12 * time.Hour).Unix()

	token, err := jwt.JWT().CreateToken(atClaims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func DelToken(id string) error {
	_, err := redis.Redis().Del(context.Background(), "jwt:"+id).Result()
	return err
}

func RefreshToken(token string, claims *jwt.CustomClaims) (string, error) {
	n, err := redis.Redis().Exists(context.Background(), "jwt:"+claims.Id).Result()
	if err != nil {
		return "", err
	}
	if n == 0 {
		return "", errors.New("您太久没有登录了，请重新登录")
	}

	newToken, err := jwt.JWT().RefreshToken(token)
	return newToken, err
}
