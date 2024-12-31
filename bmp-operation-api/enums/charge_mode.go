package enums

var ChargeMode map[int]string = map[int]string{ //创建订单sdk时候的map
	1: "CONFIG",  //按配置 postpaid_by_duration
	2: "FLOW",    //用量 postpaid_by_usage
	3: "MONTHLY", //包年包月 prepaid_by_duration
	4: "ONCE",    //一次性付费 once
}
var ChargeModeZh map[string]string = map[string]string{ //导出时候用到的。跟前端保持一直
	//"MONTHLY": "包年包月",
	//"CONFIG":  "按配置",
	"prepaid_by_duration":  "包年包月",
	"postpaid_by_duration": "按配置",
}

var ChargeModeEn map[string]string = map[string]string{ //导出时候用到的。跟前端保持一直
	//"MONTHLY": "包年包月",
	//"CONFIG":  "按配置",
	"prepaid_by_duration":  "Pay By Month",
	"postpaid_by_duration": "Pay By Configuration",
}

//创建实例sdk的map
var ChargeModeInstance map[int]string = map[int]string{
	1: "postpaid_by_duration", //按配置 postpaid_by_duration
	2: "postpaid_by_usage",    //用量 postpaid_by_usage
	3: "prepaid_by_duration",  //包年包月 prepaid_by_duration
	4: "once",                 //一次性付费 once
}
