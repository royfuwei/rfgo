package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// custom claims
type Claims struct {
	Uid                  string `json:"uid"`
	jwt.RegisteredClaims `json:",inline"`
	// *jwt.MapClaims `json:",inline"`
}

type TokenClaimsDTO struct {
	Uid string `json:"uid"`
}

type ReqJwtSign struct {
	Uid *string `json:"uid"`
}
type ReqJwtToken struct {
	Token *string `json:"token"`
}

type JwtService interface {
	JwtSign(expiresDuration time.Duration, uid *string, jwtId *string) (expiresAt int64, token string, err error)
	JwtVerify(token *string) (*jwt.MapClaims, error)
	JwtVerifyExpired(token *string) (*jwt.MapClaims, error)
	JwtDecode(token *string) (*jwt.MapClaims, error)
}

type JwtController interface {
	JwtSign(c *gin.Context)
	JwtDecode(c *gin.Context)
	JwtVerify(c *gin.Context)
	JwtVerifyExpired(c *gin.Context)
}
