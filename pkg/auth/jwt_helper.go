package auth

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rickyhuang08/gin-project/helpers"
)

type JWTHelper struct {
	Time helpers.TimeProvider
}

func NewJWTHelper(time helpers.TimeProvider) *JWTHelper {
	return &JWTHelper{
		Time: time,
	}
}

// GenerateJWT creates a signed JWT
func(m *JWTHelper) GenerateJWT(userID int, privateKey *rsa.PrivateKey) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": m.Time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}