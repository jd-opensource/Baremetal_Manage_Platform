package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/**
map或者stuct转化为json字符串
*/
func ObjToJsonString(obj interface{}) string {
	return MapToJson(InterfaceToMap(obj))
}

//任意类型转化成json，最常用
func InterfaceToJson(obj interface{}) string {
	dataType, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	dataString := string(dataType)
	return dataString
}

/**
Map对象转化为json string
*/
func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
func InterfaceToMap(obj interface{}) map[string]interface{} {

	var data = make(map[string]interface{})
	if reflect.TypeOf(obj).Kind().String() == "map" {
		d := obj.(map[string]interface{})
		for k, v := range d {
			if vv := ReflectValue2string(reflect.ValueOf(v)); vv != "" {
				data[k] = vv
			}
		}
	} else if reflect.TypeOf(obj).Kind().String() == "struct" {
		obj1 := reflect.TypeOf(obj)
		obj2 := reflect.ValueOf(obj)
		for i := 0; i < obj1.NumField(); i++ {
			if vv := ReflectValue2string(obj2.Field(i)); vv != "" {
				tagname := obj1.Field(i).Tag.Get("json")
				if tagname != "" {
					data[tagname] = vv
				} else {
					data[obj1.Field(i).Name] = vv
				}
			}
		}
	} else {
		fmt.Printf("unknown obj, type is :%s, value is:%v\n", reflect.TypeOf(obj).String(), obj)
	}
	return data

}

//ReflectValue2string 默认值都设置为"",可能会误伤，后面改的方案：1，定义成wrapper类型，2，和上下游约定字段取值不要用默认值，比如int类型取值不要出现0，bool类型不要出现false
func ReflectValue2string(obj reflect.Value) string {
	var v string
	switch obj.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if vv := strconv.FormatInt(obj.Int(), 10); vv != "0" { //默认值问题
			v = vv
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if vv := strconv.FormatUint(obj.Uint(), 10); vv != "0" { //默认值问题
			v = vv
		}
	case reflect.Float32, reflect.Float64:
		if vv := fmt.Sprintf("%f", obj.Float()); vv != "0" { //默认值问题
			v = vv
		}
	case reflect.String:
		v = obj.String()
	case reflect.Bool:
		if vv := strconv.FormatBool(obj.Bool()); vv != "false" {
			v = vv
		}
	default:
		v = ""
	}
	return v
}

//将下划线转为驼峰 name_en => nameEn 一维，二维均可
func DbToObj(obj map[string]interface{}) map[string]interface{} {
	mapObj := make(map[string]interface{})
	for kk, vv := range obj {
		if reflect.TypeOf(vv).Kind().String() == "slice" {
			//fmt.Println(vv)
			//sliceArr := []interface{}
			for k1, v1 := range vv.([]interface{}) {
				//fmt.Println(k1,v1)
				mapObj1 := make(map[string]interface{})
				for k, v := range v1.(map[string]interface{}) {
					sepStr := strings.Join(strings.Split(k, "_"), " ")
					arr := strings.Split(strings.Title(sepStr), " ")
					arr[0] = strings.ToLower(arr[0]) //hello World Haha
					mapObj1[strings.Join(arr, "")] = v
				}
				vv.([]interface{})[k1] = mapObj1
			}
			//fmt.Println(vv)
			mapObj[kk] = vv
		} else {
			sepStr := strings.Join(strings.Split(kk, "_"), " ")
			arr := strings.Split(strings.Title(sepStr), " ")
			arr[0] = strings.ToLower(arr[0])
			mapObj[strings.Join(arr, "")] = vv
		}
	}
	return mapObj
}

//将一维map的驼峰 key转化为下划线
func ObjToDb(param map[string]interface{}) map[string]interface{} {
	//str := "userNameTom" //转为user_name_tom
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`[A-Z]`)
	if reg == nil { //解释失败，返回nil
		//fmt.Println("regexp err")
		return nil
	}
	res := make(map[string]interface{})
	for key, value := range param {
		//根据规则提取关键信息
		result := reg.FindAllStringSubmatch(key, -1)
		for _, v := range result {
			//fmt.Println(v)
			key = strings.Replace(key, v[0], "_"+strings.ToLower(v[0]), 1)
		}
		res[key] = value
	}
	//fmt.Println("result1 = ", result,str)
	return res
}

func Int64ToInt(num int64) int {
	int64_num := int64(num)
	// 将 int64 转化为 int
	int_num := *(*int)(unsafe.Pointer(&int64_num))
	return int_num
}
func StringToInt64(string string) int64 {
	int64, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		fmt.Println("string to int64 error:", err.Error())
		return 0
	}
	return int64
}

/**
时间戳根据timezone解析成时间字符串
*/
func TimestampToString(timestamp int64, timezone string) string {
	if timestamp == 0 {
		return ""
	}
	t := time.Unix(timestamp, 0)
	local, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("timezone error:", timezone, err.Error())
		return ""
	}
	ts := t.In(local).String()
	fmt.Println(timestamp, timezone, ts)
	tsItems := strings.Split(ts, " ")
	if len(tsItems) >= 2 {
		ts = tsItems[0] + " " + tsItems[1]
	}
	return ts
	//if timestamp == 0 {
	//	return ""
	//}
	//return time.Unix(timestamp, 0).UTC().Format(baseLogic.DATE_UTC_FORMAT)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
