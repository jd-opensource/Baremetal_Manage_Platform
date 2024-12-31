package util

import "encoding/json"

func ObjToJson(obj interface{}) string {
	dataType, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	dataString := string(dataType)
	return dataString
}
