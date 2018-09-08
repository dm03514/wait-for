package redis

import (
	goredis "github.com/go-redis/redis"
)

type Redis struct {
	client *goredis.Client
}

func New(addr string, password string, db int) (Redis, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return Redis{
		client: client,
	}, nil
}

func (r Redis) CheckReady() (ready bool, err error) {
	status, err := r.client.Ping().Result()
	if err != nil {
		return false, err
	}

	if status == "PONG" {
		return true, nil
	}

	return false, nil
}
