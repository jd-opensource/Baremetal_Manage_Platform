package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusAlertController"] = append(beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusAlertController"],
        beego.ControllerComments{
            Method: "PrometheusAlertsReceive",
            Router: "/receiver",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusQueryController"] = append(beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusQueryController"],
        beego.ControllerComments{
            Method: "PrometheusSearchSamples",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"] = append(beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"],
        beego.ControllerComments{
            Method: "PrometheusRulesDelete",
            Router: "/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"] = append(beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"],
        beego.ControllerComments{
            Method: "PrometheusRulesList",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"] = append(beego.GlobalControllerRouter["git.jd.com/bmp-pronoea/controllers:PrometheusRulesController"],
        beego.ControllerComments{
            Method: "PrometheusRulesWrite",
            Router: "/write",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
