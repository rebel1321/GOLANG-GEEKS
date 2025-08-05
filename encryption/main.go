package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var key = "mysechfgyrneurghjiuyrtcguyretkey"

func encrypt(plainText, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	vector := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, vector)
	if err != nil {
		fmt.Println(err)
	}
	stream := cipher.NewCFBEncrypter(block, vector)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))
	fmt.Println("Cipher Text:", cipherText)
	fmt.Println("IV:", vector)
	fmt.Println("stream:", stream)
	return base64.RawStdEncoding.EncodeToString(cipherText)
}
func decrypt(cipherText, key string) string {
	block, _ := aes.NewCipher([]byte(key))
	cipherTextBytes, _ := base64.RawStdEncoding.DecodeString(cipherText)
	plainText := make([]byte, len(cipherTextBytes))
	vector := cipherTextBytes[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, vector)
	stream.XORKeyStream(plainText, cipherTextBytes[aes.BlockSize:])

	fmt.Println("Block:", block)
	fmt.Println("Cipher Text Bytes:", cipherTextBytes)
	fmt.Println("IV:", vector)
	return string(plainText)
}
func main() {
	plainText := "Hello, World!"

	cipherText := encrypt(plainText, key)
	println("Original:", plainText)
	println("Encrypted:", cipherText)
	println("Decrypted:", decrypt(cipherText, key))
}