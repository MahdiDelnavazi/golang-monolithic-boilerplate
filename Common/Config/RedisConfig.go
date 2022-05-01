package Config

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type RedisConfig struct {
	Host     string
	Password string
	Port     string
}

var Redis *redis.Client

func RedisConnection(cfg EnvironmentConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("cannot connect to redis : %s ", err)
	}
	Redis = client
	fmt.Println(pong, err)
}
