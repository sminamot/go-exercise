package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

func main() {
	plainText := "hogehoge"
	key := []byte("a1c05c11a49985dd216157277e30597c")
	iv := []byte("test-iv-12345678")

	fmt.Println("plainText:", plainText)

	encrypted, err := encrypt(plainText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("encrypted:", encrypted)

	decrypted, err := decrypt(encrypted, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("decrypted:", decrypted)
}

func encrypt(text string, key, iv []byte) (string, error) {
	data := pkcsPadding([]byte(text))

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cbc := cipher.NewCBCEncrypter(c, iv)
	encrypted := make([]byte, len(data))
	cbc.CryptBlocks(encrypted, data)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func decrypt(encrypted string, key, iv []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize || len(data)%aes.BlockSize != 0 {
		return "", errors.New("encrypted string is invalid")
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(data))
	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(decrypted, data)

	return string(pkcsUnPadding(decrypted)), nil
}

func pkcsPadding(src []byte) []byte {
	size := aes.BlockSize - (len(src) % aes.BlockSize)
	p := bytes.Repeat([]byte{byte(size)}, size)
	return append(src, p...)
}

func pkcsUnPadding(src []byte) []byte {
	p := int(src[len(src)-1])
	return src[:(len(src) - p)]
}
