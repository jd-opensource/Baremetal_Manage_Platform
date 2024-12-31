package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	helpFlag   bool
	key        string
	configFile string
	outFile    string
)

func init() {
	flag.BoolVar(&helpFlag, "h", false, "show help")
	flag.StringVar(&key, "k", "", "AES key, must 16/24/32 bytes.")
	flag.StringVar(&configFile, "f", "config.in", "config file to encrypt")
	flag.StringVar(&outFile, "o", "config.out", "output file")

}

func main() {
	flag.Parse()

	if helpFlag {
		flag.PrintDefaults()
	}

	if strings.TrimSpace(key) == "" || strings.TrimSpace(configFile) == "" || strings.TrimSpace(outFile) == "" {
		fmt.Fprintln(os.Stderr, "Specify Parameter")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if fileByte, err := ioutil.ReadFile(configFile); err != nil {
		fmt.Fprintf(os.Stderr, "Read file Error:%s\n", err.Error())
		os.Exit(2)
	} else {
		if encryptStr, e := AesEncrypt(fileByte, key); e != nil {
			fmt.Fprintf(os.Stderr, "Read file Error:%s\n", e.Error())
			os.Exit(3)
		} else {
			if errWrite := ioutil.WriteFile(outFile, []byte(encryptStr), os.ModePerm); errWrite != nil {
				fmt.Fprintf(os.Stderr, "Read file Error:%s\n", errWrite.Error())
				os.Exit(3)
			}
		}
	}

	os.Exit(0)

}

func AesEncrypt(orig []byte, key string) (string, error) {
	origData := orig
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
