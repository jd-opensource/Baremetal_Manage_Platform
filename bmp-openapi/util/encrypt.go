package util

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptEncode(pwd string) string {
	// 第二个参数是进行哈希的次数，这里采用了默认值10,数字越大生成的密码速度越慢，成本越大。但是更安全
	// bcrypt每次生成的编码是不同的，较于md5更安全
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt error:", err.Error())
		return ""
	}
	return string(bytes)
}
func BcryptVerify(pwd, hash string) bool {
	// CompareHashAndPassword 比较用户输入的明文和和数据库取出的的密码解析后是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
func GetMd5(str string, salts ...string) string {
	h := md5.New()
	h.Write([]byte(str))
	if len(salts) == 0 {
		return hex.EncodeToString(h.Sum(nil))
	} else {
		for _, salt := range salts {
			h.Write([]byte(salt))
		}
		return hex.EncodeToString(h.Sum(nil))
	}
}

func Base64Encrypt(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func Base64Decrypt(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func GetSha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}
