package errors

const (
	API_OK                    = 10000
	API_ARGUMENTS_ERROR       = 20040
	API_NO_ARGUMENTS          = 20041
	API_DB_ERROR              = 20050
	API_UNKOWN_ERROR          = 20051
	API_INTERNAL_ERROR        = 20052
	API_CALL_SERVICE_ERROR    = 20053
	API_RESOURCE_NOT_FOUND    = 20060
	API_RESOURCE_HAS_EXISTED  = 20061
	API_RESOURCE_OCCUPIED     = 20062
	API_RESOURCE_OUT_OF_RANGE = 20063
	API_RESOURCE_NOT_MATCH    = 20064
	API_LOGIN_ERROR           = 20065
	API_UPLOAD_FILE_ERROR     = 20066
	API_RESOURCE_UNAVAILABLE  = 20067
)

type codeMsg struct {
	HttpCode int
	Message  string
}

var CodeMap = make(map[int]*codeMsg)

func init() {
	CodeMap = map[int]*codeMsg{
		API_OK: {
			200,
			"OK",
		},
		API_ARGUMENTS_ERROR: {
			400,
			"参数错误",
		},
		API_NO_ARGUMENTS: {
			400,
			"参数缺失",
		},
		API_DB_ERROR: {
			500,
			"数据库错误",
		},
		API_UNKOWN_ERROR: {
			500,
			"未知错误",
		},
		API_INTERNAL_ERROR: {
			500,
			"内部错误",
		},
		API_CALL_SERVICE_ERROR: {
			500,
			"调用外部服务错误",
		},
		API_RESOURCE_NOT_FOUND: {
			400,
			"资源未找到",
		},
		API_RESOURCE_HAS_EXISTED: {
			400,
			"资源已存在",
		},
		API_RESOURCE_OCCUPIED: {
			400,
			"资源已被占用",
		},
		API_RESOURCE_OUT_OF_RANGE: {
			400,
			"资源数量超出限制",
		},
		API_RESOURCE_NOT_MATCH: {
			400,
			"资源不匹配",
		},
		API_LOGIN_ERROR: {
			401,
			"请先进行单点登录",
		},
		API_UPLOAD_FILE_ERROR: {
			400,
			"上传文件有误",
		},
		API_RESOURCE_UNAVAILABLE: {
			400,
			"资源不可用",
		},
	}
}

func GetCodeMsg(errCode int) string {
	if msg, ok := CodeMap[errCode]; ok {
		return msg.Message
	} else {
		return CodeMap[API_UNKOWN_ERROR].Message
	}
}

func GetHttpCode(errCode int) int {
	if msg, ok := CodeMap[errCode]; ok {
		return msg.HttpCode
	} else {
		return CodeMap[API_UNKOWN_ERROR].HttpCode
	}
}
