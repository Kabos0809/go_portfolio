package Config

import (
	"github.com/gomodule/redigo/redis"
)

func ConnRedis() redis.Conn {
	const p = "redis:6379"
	c, err := redis.Dial("tcp", p)
	if err != nil {	
		panic(err)
	}
	return c
}