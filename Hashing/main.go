package main

import (
	// "crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func Hashing(original string) string {
	hash := sha512.New()
	hash.Write([]byte(original))
	return base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

func main() {
	original := "Hello, World!"
	hashed := Hashing(original)
	println("Original:", original)
	println("Hashed:", hashed)
}