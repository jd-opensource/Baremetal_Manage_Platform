package util

import (
	"encoding/json"
	"math/rand"
	"runtime"
	"time"
)

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

func Obj2String(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// GetCurrentGoroutineStack 获取当前Goroutine的调用栈，便于排查panic异常
func GetCurrentGoroutineStack() string {

	const defaultStackSize = 4096

	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
