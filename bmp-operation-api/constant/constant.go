package constant

import "fmt"

var (
	OK                            RespMsg
	INVALID_ARGUMENT              RespMsg
	INVALID_ARGUMENT_WITH_ARGS    RespMsg
	NO_LOGIN_PIN                  RespMsg
	AUTH_NOT_LOGIN                RespMsg
	AUTH_BASIC_AUTH_ERROR         RespMsg
	INVALID_ARGUMENT_OUT_OF_RANGE RespMsg
	RATE_LIMIT                    RespMsg
	INTERNAL_ERROR                RespMsg
	NOT_FOUND                     RespMsg
	NOT_FOUND_WITH_ARGS           RespMsg
	CANCELLED                     RespMsg
	CANCELLEDWithArgs             RespMsg
	PermissionDenyForObject       RespMsg
	NOT_SUPPORTED                 RespMsg

	LANGUAGE_EN = "en_US"
	LANGUAGE_ZH = "zh_CN"
)

type RespMsg struct {
	Code      int    `json:"code"`
	MessageEn string `json:"message_en"`
	Messagech string `json:"message_ch"`
	Status    string `json:"status"`
}

func init() {
	OK = RespMsg{
		200,
		"OK",
		"成功",
		"OK",
	}
	INVALID_ARGUMENT = RespMsg{
		400,
		"Parameters error",
		"参数错误",
		"INVALID_ARGUMENT",
	}
	INVALID_ARGUMENT_WITH_ARGS = RespMsg{
		400,
		"%s",
		"%s",
		"INVALID_ARGUMENT",
	}
	AUTH_NOT_LOGIN = RespMsg{
		400,
		"Login Please",
		"请先登录",
		"AUTH_NOT_LOGIN",
	}
	AUTH_BASIC_AUTH_ERROR = RespMsg{
		401,
		"Unauthorized",
		"认证失败或者权限错误",
		"AUTH_ERROR_OR_PERMISSION_DENIED",
	}
	RATE_LIMIT = RespMsg{
		429,
		"Requests are too frequent",
		"请求过于频繁",
		"RATE_LIMIT",
	}
	INTERNAL_ERROR = RespMsg{
		500,
		"Internal error",
		"内部错误",
		"INTERNAL",
	}
	NOT_FOUND = RespMsg{
		404,
		"Not found",
		"找不到对象",
		"NOT_FOUND",
	}
	NOT_FOUND_WITH_ARGS = RespMsg{
		404,
		"%s Not found",
		"%s 找不到对象",
		"NOT_FOUND",
	}
	CANCELLED = RespMsg{
		400,
		"Operation cancelled",
		"取消操作",
		"CANCELLED",
	}
	NOT_SUPPORTED = RespMsg{
		400,
		"Operation not allowed",
		"不支持此操作",
		"NOT ALLOWEDS",
	}
	CANCELLEDWithArgs = RespMsg{
		400,
		"%s",
		"%s",
		"CANCELLED",
	}
	PermissionDenyForObject = RespMsg{
		401,
		"permission denied",
		"无操作权限",
		"PERMISSION DENIED",
	}
}

func BuildNotFoundWithArgs(c, e string) RespMsg {
	return RespMsg{
		404,
		fmt.Sprintf("%s Not found", e),
		fmt.Sprintf("%s 找不到对象", c),
		"NOT_FOUND",
	}
}
func BuildCANCELLEDWithArgs(c, e string) RespMsg {
	return RespMsg{
		400,
		e,
		c,
		"CANCEL",
	}
}

func BuildInvalidArgumentWithArgs(c, e string) RespMsg {
	return RespMsg{
		400,
		e,
		c,
		"INVALID_ARGUMENT",
	}
}

func BuildAuthTokenAuthError(c, e string) RespMsg {
	return RespMsg{
		401,
		e,
		c,
		"AUTH_TOKEN_AUTH_ERROR",
	}
}