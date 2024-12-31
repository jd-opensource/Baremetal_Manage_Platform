package util

import (
	"math/rand"
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
