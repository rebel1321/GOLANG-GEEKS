package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("quert12345")

func main() {
	claims := &jwt.RegisteredClaims{
		Issuer:    "testIssuer",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		Subject:  "testSubject",
		
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
	}
	fmt.Println(tokenString)
}
