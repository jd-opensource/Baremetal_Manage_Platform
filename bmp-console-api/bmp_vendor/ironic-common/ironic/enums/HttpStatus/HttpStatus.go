package HttpStatus

// refer https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/http/HttpStatus.html
const (
	ACCEPTED                        = 202 //Accepted.
	ALREADY_REPORTED                = 208 //Already Reported.
	BAD_GATEWAY                     = 502 //Bad Gateway.
	BAD_REQUEST                     = 400 //Bad Request.
	BANDWIDTH_LIMIT_EXCEEDED        = 509 //Bandwidth Limit Exceeded
	CHECKPOINT                      = 103 //Checkpoint.
	CONFLICT                        = 409 //Conflict.
	CONTINUE                        = 100 //Continue.
	CREATED                         = 201 //Created.
	EXPECTATION_FAILED              = 417 //Expectation Failed.
	FAILED_DEPENDENCY               = 424 //Failed Dependency.
	FORBIDDEN                       = 403 //Forbidden.
	FOUND                           = 302 //Found.
	GATEWAY_TIMEOUT                 = 504 //Gateway Timeout.
	GONE                            = 410 //Gone.
	HTTP_VERSION_NOT_SUPPORTED      = 505 //HTTP Version Not Supported.
	I_AM_A_TEAPOT                   = 418 //I'm a teapot.
	IM_USED                         = 226 //IM Used.
	INSUFFICIENT_STORAGE            = 507 //Insufficient Storage
	INTERNAL_SERVER_ERROR           = 500 //Internal Server Error.
	LENGTH_REQUIRED                 = 411 //Length Required.
	LOCKED                          = 423 //Locked.
	LOOP_DETECTED                   = 508 //Loop Detected
	METHOD_NOT_ALLOWED              = 405 //Method Not Allowed.
	MOVED_PERMANENTLY               = 301 //Moved Permanently. = MOVE//D_TEMPORARILY Deprecated in favor of FOUND which will be returned from HttpStatusvalueOf(302).
	MULTI_STATUS                    = 207 //Multi-Status.
	MULTIPLE_CHOICES                = 300 //Multiple Choices.
	NETWORK_AUTHENTICATION_REQUIRED = 511 //Network Authentication Required.
	NO_CONTENT                      = 204 //No Content.
	NON_AUTHORITATIVE_INFORMATION   = 203 //Non-Authoritative Information.
	NOT_ACCEPTABLE                  = 406 //Not Acceptable.
	NOT_EXTENDED                    = 510 //Not Extended
	NOT_FOUND                       = 404 //Not Found.
	NOT_IMPLEMENTED                 = 501 //Not Implemented.
	NOT_MODIFIED                    = 304 //Not Modified.
	OK                              = 200 //OK.
	PARTIAL_CONTENT                 = 206 //Partial Content.
	PAYLOAD_TOO_LARGE               = 413 //Payload Too Large.
	PAYMENT_REQUIRED                = 402 //Payment Required.
	PERMANENT_REDIRECT              = 308 //Permanent Redirect.
	PRECONDITION_FAILED             = 412 //Precondition failed.
	PRECONDITION_REQUIRED           = 428 //Precondition Required.
	PROCESSING                      = 102 //Processing.
	PROXY_AUTHENTICATION_REQUIRED   = 407 //Proxy Authentication Required.
	REQUEST_HEADER_FIELDS_TOO_LARGE = 431 //Request Header Fields Too Large.
	REQUEST_TIMEOUT                 = 408 //Request Timeout.
	REQUESTED_RANGE_NOT_SATISFIABLE = 416 //Requested Range Not Satisfiable.
	RESET_CONTENT                   = 205 //Reset Content.
	SEE_OTHER                       = 303 //See Other.
	SERVICE_UNAVAILABLE             = 503 //Service Unavailable.
	SWITCHING_PROTOCOLS             = 101 //Switching Protocols.
	TEMPORARY_REDIRECT              = 307 //Temporary Redirect.
	TOO_EARLY                       = 425 //Too Early.
	TOO_MANY_REQUESTS               = 429 //Too Many Requests.
	UNAUTHORIZED                    = 401 //Unauthorized.
	UNAVAILABLE_FOR_LEGAL_REASONS   = 451 //Unavailable For Legal Reasons.
	UNPROCESSABLE_ENTITY            = 422 //Unprocessable Entity.
	UNSUPPORTED_MEDIA_TYPE          = 415 //Unsupported Media Type.
	UPGRADE_REQUIRED                = 426 //Upgrade Required.
	URI_TOO_LONG                    = 414 //URI Too Long USE_PROXY Deprecated due to security concerns regarding in-band configuration of a proxy
	VARIANT_ALSO_NEGOTIATES         = 506 //Variant Also Negotiates
)
