package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret = []byte("MySignature")

type AuthClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"userId"`
}

func CreateToken(userId uuid.UUID) (string, error) {
	// claims := &jwt.RegisteredClaims{
	// 	Issuer:    "auth-claims",
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// }
	claims := &AuthClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-claims",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	ss, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func Valid(ss string) {
	tok, err := jwt.ParseWithClaims(
		ss,
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token: %v\n", tok)
	fmt.Printf("token valid: %t\n", tok.Valid)
}
