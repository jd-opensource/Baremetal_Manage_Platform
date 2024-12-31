package util

func TurnInterfaceToInt64(org interface{}) int64 {
	var data int64 = 0
	switch org.(type) {
	case int64:
		data = org.(int64)
	case int:
		data = int64(org.(int))
	case float64:
		data = int64(org.(float64))
	}

	return data
}

func TurnInterfaceToString(org interface{}) string {
	var str string = ""
	switch org.(type) {
	case string:
		str = org.(string)
	}

	return str
}
