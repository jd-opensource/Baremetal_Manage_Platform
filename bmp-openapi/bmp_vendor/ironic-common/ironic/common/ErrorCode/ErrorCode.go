package ErrorCode

const (
	// 此操作需要的预置条件不满足 400
	FAILED_PRECONDITION = "FAILED_PRECONDITION"

	// 参数错误 400
	INVALID_ARGUMENT = "INVALID_ARGUMENT"

	// 缺少参数 400
	MISSING_ARGUMENT = "MISSING_ARGUMENT"

	// 参数类型不匹配400
	MISMATCH_ARGUMENT_TYPE = "MISMATCH_ARGUMENT_TYPE"

	// 无效的action 400
	INVALID_ACTION = "INVALID_ACTION"

	// 参数超出范围 400
	OUT_OF_RANGE = "OUT_OF_RANGE"

	// 缺少Multipart 400
	MISSING_MULTIPART = "MISSING_MULTIPART"

	// 缺少header信息 400
	MISSING_HEADER = "MISSING_HEADER"

	// pathVariable参数错误 400
	INVALID_PATHVARIABLE = "INVALID_PATHVARIABLE"

	// 验证失败 401
	UNAUTHENTICATED = "UNAUTHENTICATED"

	// 用户未激活 402
	USER_NOT_ACTIVE = "USER_NOT_ACTIVE"

	// 没有权限 403
	HTTP_FORBIDDEN = "HTTP_FORBIDDEN"

	// 找不到对象 404
	NOT_FOUND = "NOT_FOUND"

	// 操作终止 409
	ABORTED = "ABORTED"

	// 对象已存在 409
	ALREADY_EXISTS = "ALREADY_EXISTS"

	// 媒体类型不支持 415
	UNSUPPORTED_MEDIA_TYPE = "UNSUPPORTED_MEDIA_TYPE"

	// 不支持的请求方法 405
	UNSUPPORTED_HTTP_REQUEST_METHOD = "UNSUPPORTED_HTTP_REQUEST_METHOD"

	// 配额不足 429
	QUOTA_EXCEEDED = "QUOTA_EXCEEDED"

	// 库存不足429
	STOCK_SHORTAGE = "STOCK_SHORTAGE"

	// 未知错误 500
	UNKNOWN = "UNKNOWN"

	// 内部错误 500
	INTERNAL = "INTERNAL"

	// 服务不可用 503
	UNAVAILABLE = "UNAVAILABLE"

	// 会话网关 Timeout 504
	SESSION_GATEWAY_TIMEOUT = "SESSION_GATEWAY_TIMEOUT"

	// 参数错误
	PARAMETER_ERROR = "PARAMETER_ERROR"

	//com.jcloud.cps.ironic.common.ErrorCode
	UNSUPPORTED_OPERATION = "UNSUPPORTED_OPERATION"

	// 子网没有可用的内网IP可以分配 400
	EXHAUSTED_SUBNET = "EXHAUSTED_SUBNET"

	EXHAUSTED_PUBLIC_IP = "EXHAUSTED_PUBLIC_IP"

	UNSUPPORTED_RAID_TYPE = "UNSUPPORTED_RAID_TYPE"

	CHANGE_ARE_NOT_ALLOWED = "CHANGE_ARE_NOT_ALLOWED"

	// 权限校验失败
	NO_PERMISSION = "NO_PERMISSION"

	// 资源被占用
	RESOURCE_IN_USE = "RESOURCE_IN_USE"

	// 内网IP超出范围
	PRIVATE_IP_NOT_IN_SUBNET = "PRIVATE_IP_NOT_IN_SUBNET"

	// EIP已经绑定
	EIP_ALREADY_ASSOCIATE = "EIP_ALREADY_ASSOCIATE"

	// EIP未绑定
	EIP_HAS_NOT_ASSOCIATE = "EIP_HAS_NOT_ASSOCIATE"

	// 子网未开通IPv6
	SUBNET_HAS_NOT_ASSIGN_IPV6 = "SUBNET_HAS_NOT_ASSIGN_IPV6"

	// IPV6已经被分配
	IPV6_ALREADY_ASSIGN = "IPV6_ALREADY_ASSIGN"

	// 实例已经绑定EIP
	INSTANCE_ALREADY_ASSOCIATE_EIP = "INSTANCE_ALREADY_ASSOCIATE_EIP"

	// 实例还没有绑定EIP
	INSTANCE_HAS_NOT_ASSOCIATE_EIP = "INSTANCE_HAS_NOT_ASSOCIATE_EIP"

	// 实例已经绑定IPV6
	INSTANCE_ALREADY_ASSOCIATE_IPV6 = "INSTANCE_ALREADY_ASSOCIATE_IPV6"

	// 基础网路不支持EIP
	UNSUPPORTED_EIP_IN_BASIC_NETWORK = "UNSUPPORTED_EIP_IN_BASIC_NETWORK"

	// VPC网络不支持调整带宽
	UNSUPPORTED_MODIFY_BANDWIDTH_IN_VPC_NETWORK = "UNSUPPORTED_MODIFY_BANDWIDTH_IN_VPC_NETWORK"

	// 不支持批量指定内网IP
	UNSUPPORTED_BATCH_SPECIFYING_PRIVATE_IP = "UNSUPPORTED_BATCH_SPECIFYING_PRIVATE_IP"

	// 不支持批量指定别名IP
	UNSUPPORTED_BATCH_SPECIFYING_ALIAS_IP = "UNSUPPORTED_BATCH_SPECIFYING_ALIAS_IP"

	// 密钥对正在被实例使用
	KEYPAIR_IS_USING_BY_INSTANCE = "KEYPAIR_IS_USING_BY_INSTANCE"

	ARGUMENT_ERROR = "ARGUMENT_ERROR"
)
