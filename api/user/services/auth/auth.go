package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

var signingKey = []byte("eeae315eb1b54ac8cd1cc67a32087073e8cd4b022dd75f40eb553cde9039cd80")

type CustomClaims struct {
	Id uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uuid.UUID) (string, error) {
	claims := CustomClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(token string) {

}
