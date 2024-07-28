package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// custom claims
type Claims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

type TokenClaims interface{}

type JwtService interface {
	JwtSign(expiresTime time.Duration, uid *string, jwtId *string) (expiresAt int64, token string, err error)
}
