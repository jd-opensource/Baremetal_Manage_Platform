package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/gofrs/uuid"
)

func GenUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

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

//生成8-30位随机密码
func GenerateRandPassWord() string {
	a := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	d := "0123456789()`~!@#$%&*_-+=|{}[]:\";'<>,.?/"
	c := strings.Split(d, "")
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(a)) //生成0-数组长度之间的随机整数
	y := rand.Intn(len(b)) //生成0-数组长度之间的随机整数
	z := rand.Intn(len(c)) //生成0-数组长度之间的随机整数
	//fmt.Println(c)
	s := make([]string, 0)
	s = append(s, a[x])
	s = append(s, b[y])
	s = append(s, c[z])
	//以上生成了3个固定长度的字符串，这是必须要有的3个
	var arr = append(append(a, b...), c...)

	for i := 1; i <= 5; i++ {
		rand := rand.Intn(len(arr)) //生成0-数组长度之间的随机整数
		s = append(s, arr[rand])
	}
	//以上组成了至少8位的一个随机字符，下面生成9-30之间的随机长度的字符串
	length := rand.Intn(23) //生成0-22之间的数，前闭区间，后开区间，跟前面的8位拼接，组成8-30之间的随机数
	if length != 0 {
		for i := 1; i <= length; i++ {
			rand := rand.Intn(len(arr)) //生成0-数组长度之间的随机整数
			s = append(s, arr[rand])
		}
	}
	str := "" //下面将最终的数组打乱，生成新的字符串
	for i := 0; i < len(s); i++ {
		rand := rand.Intn(len(s))
		str += s[rand]
	}
	return str
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
