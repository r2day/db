package db

import (
	"tbkt/logger"
	"time"

	"github.com/redis/go-redis/v9"

	log "github.com/sirupsen/logrus"
)

// RDB 定义redis链接池
var RDB *redis.Client

// InitRedisDB 初始化redis链接池
func InitRedisDB(redisAddr string, db int, poolSize int) {
	RDB = redis.NewClient(&redis.Options{
		Addr:        redisAddr,        // Redis地址
		Password:    "",               // Redis账号
		DB:          db,               // Redis库
		PoolSize:    poolSize,         // Redis连接池大小
		MaxRetries:  3,                // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
	})
	pong, err := RDB.Ping().Result()
	if err == redis.Nil {
		logger.Info("Redis异常")
		log.WithField("redisAddr", redisAddr).
			WithField("db", db).
			WithField("poolSize", poolSize).Error(err)
	} else if err != nil {
		logger.Info("Redis异常")
		log.WithField("redisAddr", redisAddr).
			WithField("db", db).
			WithField("poolSize", poolSize).Error(err)
	} else {
		logger.Info("Redis异常")
		log.WithField("redisAddr", redisAddr).
			WithField("db", db).
			WithField("poolSize", poolSize).Info("connect redis successful")
	}
}
