package config

import (
	"fmt"
	"os"
)

type Redis struct {
	addr     string
	Port     int
	Password string
}

func ReadRedisConfig() *Redis {
	return &Redis{
		addr:     os.Getenv("REDIS_ADDR"),
		Port:     6379,
		Password: os.Getenv("REDIS_PASS"),
	}
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.addr, r.Port)
}
