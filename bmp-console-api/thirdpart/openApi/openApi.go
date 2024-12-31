package openApi

import (
	"fmt"

	sdk "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client"
	beego "github.com/beego/beego/v2/server/web"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var (
	SdkClient *sdk.CPS
)

func Init() {

	host, _ := beego.AppConfig.String("openapi_host")
	port, _ := beego.AppConfig.String("openapi_port")
	host = fmt.Sprintf("%s:%s", host, port)
	fmt.Println("openapi host:", host)

	token, _ := beego.AppConfig.String("openapi_token")
	/*
		cfg := sdk.TransportConfig{
			Host:     host,
			BasePath: "/v1",
			Schemes:  []string{"http"},
		}
		SdkClient = sdk.NewHTTPClientWithConfig(nil, &cfg)
	*/

	r := httptransport.New(host, "/v1", []string{"http"})
	r.SetDebug(true)
	r.DefaultAuthentication = httptransport.BearerToken(token)
	/*
	  r.DefaultMediaType = runtime.XMLMime
	  r.Consumers = map[string]runtime.Consumer{
	    runtime.XMLMime: runtime.XMLConsumer(),
	  }
	  r.Producers = map[string]runtime.Producer{
	    "application/xhtml+xml": runtime.XMLProducer(),
	  }
	*/
	SdkClient = sdk.New(r, strfmt.Default)

}
