package enums

var OrderType map[int]string = map[int]string{
	1: "NEW",            //新购
	2: "RENEW",          //续费
	3: "RESIZE_FORMULA", //配置变更
}
