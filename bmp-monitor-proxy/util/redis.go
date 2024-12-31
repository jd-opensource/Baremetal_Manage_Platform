package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var redisMonitorNamespace string

var redisOnce sync.Once

type redisUtil struct {
	*redis.Client
}

// RedisUtil singleton redis util
var RedisUtil *redisUtil

// InitRedis initial redisUtil instance
func InitRedis(namespace, idc, addr, passwd string, db int) error {
	redisOnce.Do(func() {
		if RedisUtil == nil {

			RedisUtil = &redisUtil{
				redis.NewClient(&redis.Options{
					Addr:     addr,
					Password: passwd, // no password set
					DB:       db,     // use default DB
				}),
			}

			redisMonitorNamespace = namespace
			if idc != "" {
				redisMonitorNamespace = namespace + ":" + idc
			}

		}
	})

	return RedisUtil.Ping().Err()
}

// SetObjectToRedisWithExpire set string object with a expiration
func SetObjectToRedisWithExpire(key, value string, sec int) error {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	err := RedisUtil.Set(fullKey, value, time.Second*time.Duration(sec)).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetObjectToRedisWithExpire set string object with a expiration
func SetObjectWithExpire(key string, sec int) error {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	err := RedisUtil.Expire(fullKey, time.Second*time.Duration(sec)).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetObjectToRedis set object to Redis
func SetObjectToRedis(key string, value string) error {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}
	fmt.Println("redis set full key is:" + fullKey)
	fmt.Println("redis set value is:" + value)
	err := RedisUtil.Set(fullKey, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetObjectFromRedis Get object from Redis
func GetObjectFromRedis(key string) (string, error) {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	ret, err := RedisUtil.Get(fullKey).Result()
	if err != nil {
		return "", err
	}

	return ret, nil
}

func DelObjectFromRedis(keys []string) error {

	if len(keys) <= 0 {
		return nil
	}

	var fullKeys []string
	for _, v := range keys {
		if v != "" {
			fullKeys = append(fullKeys, redisMonitorNamespace+":"+v)
		}
	}

	if len(fullKeys) > 0 {
		err := RedisUtil.Del(fullKeys...).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// HSetObjectToRedis hash set object to Redis
func HSetObjectToRedis(key, field, value string) error {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	err := RedisUtil.HSet(fullKey, field, value).Err()
	if err != nil {
		return err
	}

	return nil
}

// HSetObjectToRedis hash set object to Redis
func HMSetObjectToRedis(key string, data map[string]interface{}) error {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	err := RedisUtil.HMSet(fullKey, data).Err()
	if err != nil {
		return err
	}

	return nil
}

// HGetObjectFromRedis hash get from redis
func HGetAllFromRedis(key string) (map[string]string, error) {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	ret, err := RedisUtil.HGetAll(fullKey).Result()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// HGetObjectFromRedis hash get from redis
func HGetObjectFromRedis(key, field string) (string, error) {

	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	ret, err := RedisUtil.HGet(fullKey, field).Result()
	if err != nil {
		return "", err
	}
	return ret, nil
}

func HDelObjectFromRedis(key string, fields []string) error {
	fullKey := redisMonitorNamespace
	if key != "" {
		fullKey = redisMonitorNamespace + ":" + key
	}

	if len(fields) <= 0 {
		return nil
	}

	err := RedisUtil.HDel(fullKey, fields...).Err()

	if err != nil {
		return err
	}
	return nil
}
