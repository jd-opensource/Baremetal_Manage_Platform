package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//填充/裁剪字符串
func Padding(source string, target_len int, ahead bool, padding_char rune) string {
	r := []rune(source)
	if len(r) > target_len { //裁剪
		if ahead { //保留左边
			return string(r[:target_len])
		} else { //保留右边
			return string(r[len(r)-target_len:])
		}

	} else if len(r) < target_len { //填充
		pads := []rune{}
		for i := 0; i < target_len-len(r); i++ {
			pads = append(pads, padding_char)
		}
		if ahead { //左向填充
			return string(append(pads, r...))
		} else { //右向填充
			return string(append(r, pads...))
		}
	}
	return source

}

func String2int8(v string) int8 {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("String2int8 error:%s", err.Error()))
	}
	return int8(i)
}

func String2int(v string) int {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("String2int8 error:%s", err.Error()))
	}
	return int(i)
}

func String2Bool(v string) bool {
	if strings.ToLower(v) == "true" {
		return true
	}
	return false
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
