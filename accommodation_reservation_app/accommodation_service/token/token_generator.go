package token

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Role string

const (
	Host  Role = "HOST"
	Guest Role = "GUEST"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Role       Role
	Uid        string
	jwt.StandardClaims
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The Token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}
