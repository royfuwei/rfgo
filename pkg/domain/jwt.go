package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// custom claims
type Claims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

type TokenClaims interface{}

type ReqJwtSign struct {
	Uid *string `json:"uid"`
}
type ReqJwtDecode struct {
	Token *string `json:"token"`
}

type JwtService interface {
	JwtSign(expiresTime time.Duration, uid *string, jwtId *string) (expiresAt int64, token string, err error)
	JwtVerify(token *string) (*Claims, error)
	JwtDecode(token *string) (TokenClaims, error)
}

type JwtController interface {
	JwtSign(c *gin.Context)
	JwtDecode(c *gin.Context)
}
