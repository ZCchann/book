package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type rdb struct {
	*redis.Client
}

var r = new(rdb)

func Redis() *rdb {
	return r
}
func (r *rdb) Connect(addr, password string, db int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	r.Client = client
	return nil
}
