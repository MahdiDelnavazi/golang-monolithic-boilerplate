package Config

import (
	"github.com/go-redis/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stretchr/testify/require"
	"testing"
)

type RedisConfigTest struct {
	Host     string `env:"REDIS_HOST" env-default:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD" env-default:""`
	Port     string `env:"REDIS_PORT" env-default:"6379"`
}

func TestRedisConnection(t *testing.T) {
	config := RedisConfigTest{}
	err := cleanenv.ReadConfig("../../.test.env", &config)
	require.NoError(t, err)
	require.NotNil(t, config)

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       0,
	})
	require.NotNil(t, client)
	require.NotEmpty(t, config)

	pong, err := client.Ping().Result()
	require.NoError(t, err)
	require.NotNil(t, pong)
	require.NotEmpty(t, pong)

}
