package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	AES_KEY = "ABCDEFGHIJKLMNOP"
	AES_IV  = "0123456789ABCDEF"
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

func Base64Encrypt(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func Base64Decrypt(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func GetSha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}

func AesEncrypt(data []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher([]byte(AES_KEY))
	if err != nil {
		println(err.Error())
		return nil, err
	}

	content := PKCS7Padding2(data, aesBlockEncrypt.BlockSize())
	cipherBytes := make([]byte, len(content))
	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, []byte(AES_IV))
	aesEncrypt.CryptBlocks(cipherBytes, content)
	return cipherBytes, nil
}

func AesDecrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err = aes.NewCipher([]byte(AES_KEY))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, []byte(AES_IV))
	aesDecrypt.CryptBlocks(decrypted, src)
	return PKCS7UnPadding2(decrypted), nil
}

//PKCS7Padding 使用PKCS7进行填充
func PKCS7Padding2(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//PKCS7UnPadding 解PKCS7填充
func PKCS7UnPadding2(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	if length < unPadding {
		fmt.Println(time.Now().String(), "PKCS7UnPadding error")
		return nil
	}
	return origData[:(length - unPadding)]
}
