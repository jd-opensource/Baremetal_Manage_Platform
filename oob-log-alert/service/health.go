package service

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/util"
)

// RedisCheck beego health check struct, implement Check method
type RedisCheck struct {
}

// Check health check for redis
func (redis *RedisCheck) Check() error {
	err := redis.ping()
	if err != nil {
		return err
	}
	return nil

}

func (redis *RedisCheck) ping() error {
	return util.RedisUtil.Ping().Err()

}

type MySQLCheck struct {
}

// Check health check for MySQL
func (redis *MySQLCheck) Check() error {

	return nil

}
