package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 10 * time.Second,
		MaxActive:   300,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp", fmt.Sprintf(
					"%s:%s",
					os.Getenv("REDIS_HOST"),
					os.Getenv("REDIS_PORT"),
				),
			)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
