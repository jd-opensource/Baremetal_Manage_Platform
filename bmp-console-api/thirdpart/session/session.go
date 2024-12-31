package session

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

func Init() error {

	cookieName, _ := beego.AppConfig.String("sessionname")
	sessionLife, _ := beego.AppConfig.Int64("sessionlife")

	redisHost, _ := beego.AppConfig.String("bmp_redis_host")
	redisPort, _ := beego.AppConfig.Int("bmp_redis_port")
	redisPassword, _ := beego.AppConfig.String("bmp_redis_password")
	redisProvider := fmt.Sprintf("%s:%d,10,%s", redisHost, redisPort, redisPassword)
	fmt.Println("session.redis.provicer:", redisProvider)

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = cookieName
	// beego.BConfig.WebConfig.Session.SessionCookieLifeTime = sessionLife
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = sessionLife
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = redisProvider

	return nil
}
