package util

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

func ToString(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%d", val)
	case reflect.String:
		return val.(string)
	case reflect.Array, reflect.Slice, reflect.Ptr, reflect.Interface, reflect.Map, reflect.Struct:
		if s, err := json.Marshal(val); err == nil {
			return string(s)
		} else {
			log.Print(fmt.Sprintf("val for %+v ToString error: %s", val, err.Error()))
			return ""
		}
	case reflect.Func, reflect.Chan, reflect.UnsafePointer:
		log.Print(fmt.Sprintf("val for %+v ToString error", val))
		return ""
	}
	return ""
}

func Time2String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//GetRandString 借鉴RandomStringUtils.random包生成随机字符串
// length 随机字符串的长度(不包括前缀)
// uletter 随机字符串是否包含大写字母字符
// dletter 随机字符串是否包含小写字母字符
// number 随机字符串是否包含数字字符
func GetRandString(prefix string, length int32, uletter, dletter, number bool) string {
	rand.Seed(time.Now().UnixNano())
	uletters := []rune("ABCDEFGHIGKLMNOPQRSTUVWXYZ")
	dletters := []rune("abcdefghijklmnopqrstuvwxyz")
	numbers := []rune("0123456789")
	source := []rune{}
	if uletter {
		source = append(source, uletters...)
	}
	if dletter {
		source = append(source, dletters...)
	}
	if number {
		source = append(source, numbers...)
	}
	if len(source) == 0 {
		return "atLeastContainsLetterOrNumber"
	}
	b := make([]rune, length)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}
	return prefix + string(b)
}

// 跟现有环境一致，格式：aafc4d1b-f696-4ec2-924f-6fd012aa8345
func GenerateRandUuid() string {
	s := GetRandString("", 32, false, true, true)
	runes := []rune(s)
	return string(runes[0:8]) + "-" + string(runes[8:12]) + "-" + string(runes[12:16]) + "-" + string(runes[16:20]) + "-" + string(runes[20:32])
}

func GetUuid(pre string) string {
	rand.Seed(time.Now().UnixNano())
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < 28; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return pre + string(result)
}
