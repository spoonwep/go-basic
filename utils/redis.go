package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var Ctx = context.TODO()
var RedisClient *redis.Client

type KeyResult struct {
	Key       string `json:"key"`
	IsSuccess bool   `json:"is_success"`
}

func InitRedis() {
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     os.Getenv("REDIS_PASS"),
		DB:           redisDb,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		logrus.Fatalf("redis connect error: host: %s, detail: %#v", redisAddr, err)
	}
}

func RedisGet(key string) (interface{}, error) {
	return RedisClient.Get(Ctx, key).Result()
}

func RedisSet(key, value string, timeout time.Duration) error {
	return RedisClient.Set(Ctx, key, value, timeout).Err()
}

func RedisDel(key string) {
	RedisClient.Del(Ctx, key)
}

func RedisLPop(key string) (string, error) {
	return RedisClient.LPop(Ctx, key).Result()
}

func RedisLPush(key, value string) {
	RedisClient.LPush(Ctx, key, value)
}
