package main

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims 사용자 정의 클레임 구조체
type MyCustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// parseToken 토큰 검증 및 파싱 함수
func parseToken(tokenStr string, secretKey []byte) (*MyCustomClaims, error) {
	// 토큰 파싱 및 검증
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	// 비밀키
	secretKey := []byte("your-256-bit-secret")

	// 예제 토큰 (실제 토큰 사용 필요)
	tokenString := "your.jwt.token"

	// 토큰 검증
	claims, err := parseToken(tokenString, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UserID: %s\n", claims.UserID)
}
