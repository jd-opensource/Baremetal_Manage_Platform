package redis

import (
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
func InitRedis(namespace, addr, passwd string, db int) error {
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

		}
	})
	return RedisUtil.Ping().Err()
}

// SetObjectToRedisWithExpire set string object with a expiration
func SetObjectToRedisWithExpire(key, value string, sec int) error {

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}

	err := RedisUtil.Set(fullKey, value, time.Second*time.Duration(sec)).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetObjectToRedis set object to Redis
func SetObjectToRedis(key string, value string) error {

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}

	err := RedisUtil.Set(fullKey, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// SetObjectToRedis set object to Redis
func SetNxObjectToRedis(key string, value string, expire int) (bool, error) {

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}
	return RedisUtil.SetNX(fullKey, value, time.Second*time.Duration(expire)).Result()
}

func Keys(key string) ([]string, error) {
	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}
	return RedisUtil.Keys(fullKey).Result()
}

// GetObjectFromRedis Get object from Redis
func GetObjectFromRedis(key string) (string, error) {

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
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
			var fullKey string
			if redisMonitorNamespace == "" {
				fullKey = v
			} else {
				fullKey = redisMonitorNamespace
				if v != "" {
					fullKey = redisMonitorNamespace + ":" + v
				}
			}
			fullKeys = append(fullKeys, fullKey)
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

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}

	err := RedisUtil.HSet(fullKey, field, value).Err()
	if err != nil {
		return err
	}

	return nil
}

// HGetObjectFromRedis hash get from redis
func HGetObjectFromRedis(key, field string) (string, error) {

	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}

	ret, err := RedisUtil.HGet(fullKey, field).Result()
	if err != nil {
		return "", err
	}
	return ret, nil
}

func HDelObjectFromRedis(key string, fields []string) error {
	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
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

func SaddObjectToRedis(key string, members []interface{}) error {
	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}
	if len(members) <= 0 {
		return nil
	}

	err := RedisUtil.SAdd(fullKey, members...).Err()

	if err != nil {
		return err
	}
	return nil
}

func SremObjectFromRedis(key string, members []interface{}) error {
	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}
	if len(members) <= 0 {
		return nil
	}

	err := RedisUtil.SRem(fullKey, members...).Err()

	if err != nil {
		return err
	}
	return nil
}

func SmembersFromRedis(key string) ([]string, error) {
	var fullKey string
	if redisMonitorNamespace == "" {
		fullKey = key
	} else {
		fullKey = redisMonitorNamespace
		if key != "" {
			fullKey = redisMonitorNamespace + ":" + key
		}
	}

	members, err := RedisUtil.SMembers(fullKey).Result()

	if err != nil {
		return nil, err
	}
	return members, nil
}
