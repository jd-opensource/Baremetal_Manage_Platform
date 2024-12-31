package router

import (
	"coding.jd.com/bmp/agent-proxy-jdstack/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	// api namespace router
	ns :=
		beego.NewNamespace("/api/v1/agent",
			beego.NSRouter("/put", &controllers.ProxyController{}, "post:Put"),
			beego.NSRouter("/heartbeat", &controllers.ProxyController{}, "post:Heartbeat"),
		)

	beego.AddNamespace(ns)

	apiNS :=
		beego.NewNamespace("/api/v1/",
			beego.NSRouter("/describeAgentStatus", &controllers.StatusController{}, "post:DescribeAgentStatus"),
		)
	beego.AddNamespace(apiNS)

	// api namespace router
	tagsNs :=
		beego.NewNamespace("/api/v1/tags",
			beego.NSRouter("/", &controllers.DeviceTagController{}, "post:GetDeviceTags"),
		)

	beego.AddNamespace(tagsNs)

}
