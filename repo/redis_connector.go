package repo

import (
	"encoding/json"
	"fmt"
	"member_service_frame/config"

	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func newRedisConfig() *redisConfig {
	var set = config.Setting.GetRedisSetting()
	config := new(redisConfig)
	j, _ := json.Marshal(set)
	json.Unmarshal(j, config)
	return config
}

// GetRedisConnecter returns a Redis client connected to the specified Redis pool.
func GetRedisConnecter(poolNum int) *redis.Client {
	var redisSetting *redisConfig = newRedisConfig()
	var client *redis.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisSetting.Host, redisSetting.Port),
		Password: redisSetting.Password,
		DB:       poolNum,
	})
	return client
}
