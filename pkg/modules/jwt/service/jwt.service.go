package service

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/glog"
	"github.com/royfuwei/rfgo/pkg/domain"
)

type jwtService struct {
	privateKeyPath string
	publicKeyPath  string
	verifyKey      *rsa.PublicKey  // openssl genpkey -algorithm RSA -out jwt.rsa -pkeyopt rsa_keygen_bits:2048
	signKey        *rsa.PrivateKey // openssl rsa -in jwt.rsa -pubout > jwt.rsa.pub
}

func NewJwtService(privateKeyPath, publicKeyPath string) domain.JwtService {
	svc := &jwtService{
		privateKeyPath: privateKeyPath,
		publicKeyPath:  publicKeyPath,
	}
	svc.setRsaKeys()
	return svc
}

func (svc *jwtService) JwtSign(expiresDuration time.Duration, uid *string, jwtId *string) (expiresAt int64, token string, err error) {
	now := time.Now()
	expiresTime := jwt.NewNumericDate(now.Add(expiresDuration * time.Second))
	expiresAt = expiresTime.Time.Unix()
	claims := domain.Claims{
		Uid: *uid,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: expiresTime,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "royfuwei/rfgo",
			Subject:   "royfuwei",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err = tokenClaims.SignedString(svc.signKey)
	return expiresAt, token, err
}

func (svc *jwtService) JwtDecode(token *string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	parser := jwt.NewParser()
	_, _, err := parser.ParseUnverified(*token, claims)
	return claims, err
}

func (svc *jwtService) JwtVerify(token *string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	parser := jwt.NewParser()
	_, err := parser.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return svc.verifyKey, nil
	})
	return claims, err
}

func (svc *jwtService) JwtVerifyExpired(token *string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	parser := jwt.NewParser(jwt.WithExpirationRequired())
	parseToken, err := parser.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return svc.verifyKey, nil
	})
	isErr := false
	if err != nil {
		switch {
		case parseToken.Valid:
			fmt.Println("You look nice today")
			isErr = true
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
			isErr = true
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
			isErr = true
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			// fmt.Println("Timing is everything")
			isErr = false
		default:
			fmt.Println("Couldn't handle this token:", err)
			isErr = true
		}
	}
	if claims, ok := parseToken.Claims.(*jwt.MapClaims); ok && !isErr {
		return claims, nil
	} else {
		return nil, err
	}
}

func (svc *jwtService) setRsaKeys() {
	signBytes, err := os.ReadFile(svc.privateKeyPath)
	svc.fatal(err)
	verifyBytes, err := os.ReadFile(svc.publicKeyPath)
	svc.fatal(err)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	svc.fatal(err)
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	svc.fatal(err)
	svc.verifyKey = verifyKey
	svc.signKey = signKey
}

func (svc *jwtService) fatal(err error) {
	if err != nil {
		glog.Fatal(err)
	}
}
