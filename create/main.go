package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims 사용자 정의 클레임 구조체
type MyCustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// createToken JWT 생성 함수
func createToken(secretKey []byte) string {
	// 사용자 정의 클레임 설정
	claims := MyCustomClaims{
		UserID: "1234567890",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "test",
		},
	}

	// 새 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 토큰 서명
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func main() {
	// 비밀키 (실제 사용 시 안전하게 보관)
	secretKey := []byte("your-256-bit-secret")

	// JWT 생성
	tokenString := createToken(secretKey)
	fmt.Println("Token:", tokenString)
}
