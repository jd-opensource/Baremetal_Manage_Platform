package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncrypt AES encrypt method. key must 128/192/256 bits.
//
// example:
//   orig := "hello world"
//   key := "12345678123456781234567812345678"
//   fmt.Println("原文：", orig)
//   encryptCode,_ := AesEncrypt(orig, key)
//   fmt.Println("密文：" , encryptCode)
func AesEncrypt(orig string, key string) (string, error) {
	origData := []byte(orig)
	k := []byte(key)
	if block, err := aes.NewCipher(k); err != nil {
		return "", err
	} else {
		blockSize := block.BlockSize()
		origData = pkcs7Padding(origData, blockSize)
		blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
		cryted := make([]byte, len(origData))
		blockMode.CryptBlocks(cryted, origData)
		return base64.StdEncoding.EncodeToString(cryted), nil
	}
}

// AedDecrypt AES decrypt.key must 128/192/256 bits.
//
// example:
//  key := "12345678123456781234567812345678"
//  data := bpIjQ4yVwCt0RMI0JDOfTA==
//  decryptCode,_ := AesDecrypt(data, key)
//  fmt.Println("解密结果：", decryptCode)
func AesDecrypt(cryted string, key string) (string, error) {
	if crytedByte, err := base64.StdEncoding.DecodeString(cryted); err != nil {
		return "", err
	} else {
		k := []byte(key)
		if block, e := aes.NewCipher(k); e != nil {
			return "", e
		} else {
			blockSize := block.BlockSize()
			blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
			orig := make([]byte, len(crytedByte))
			blockMode.CryptBlocks(orig, crytedByte)
			orig = pkcs7UnPadding(orig)
			return string(orig), nil
		}
	}
}

func pkcs7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
