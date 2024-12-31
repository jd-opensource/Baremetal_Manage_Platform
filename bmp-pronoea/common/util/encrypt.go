package util

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

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

func GetUuid() string {
	return fmt.Sprintf("%v", uuid.NewV4())
}

func GetSha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}
