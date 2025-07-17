package core

import (
	"blogx_server/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	r := global.Config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		logrus.Fatalf("redis connect err: %v", err)
	}
	logrus.Info("redis connect success")
	return redisClient
}
