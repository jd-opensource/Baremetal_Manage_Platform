package util

import (
	"fmt"
	"strconv"
	"strings"
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
