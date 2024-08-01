package service

import (
	"crypto/rsa"
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

func (svc *jwtService) JwtSign(expiresTime time.Duration, uid *string, jwtId *string) (expiresAt int64, token string, err error) {
	now := time.Now()
	expiresAt = now.Add(time.Duration(time.Duration(3600).Seconds())).Unix()
	claims := jwt.MapClaims{
		"uid": uid,
		"iss": "rfgo",
		"iat": now.Unix(),
		"exp": expiresAt,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err = tokenClaims.SignedString(svc.signKey)
	return expiresAt, token, err
}

func (svc *jwtService) JwtDecode(token *string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return svc.verifyKey, nil
	}, jwt.WithStrictDecoding())
	return claims, err
}

func (svc *jwtService) JwtVerify(token *string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return svc.verifyKey, nil
	})
	return claims, err
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
