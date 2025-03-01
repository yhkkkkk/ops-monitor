package client

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"ops-monitor/internal/global"
)

var Redis *redis.Client

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Pass,
		DB:       0, // 使用默认的数据库
	})

	// 尝试连接到 Redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("redis Connection Failed %s", err)
		return nil
	}

	Redis = client

	return client

}
