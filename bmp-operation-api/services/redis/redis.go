package redis

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis"
)

var RedisCli *redis.Client

// InitRedis initial redisUtil instance
func Init() error {
	var host, port, password string
	var dbNum int
	//namespace, _ = beego.AppConfig.String("bmp_redis_namespace")
	host, _ = beego.AppConfig.String("bmp_redis_host")
	port, _ = beego.AppConfig.String("bmp_redis_port")
	password, _ = beego.AppConfig.String("bmp_redis_password")
	dbNum = beego.AppConfig.DefaultInt("bmp_redis_db", 0)
	//addr := fmt.Sprintf("%s:%s", host, port)
	//return redisInit.InitRedis(namespace, addr, password, dbNum)

	RedisCli = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       dbNum,    // use default DB
	})
	_, err := RedisCli.Ping().Result()
	if err != nil {
		fmt.Println("redis connect error:", RedisCli, err)
		return err
	}
	return nil
}
