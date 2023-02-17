package app

import (
	"giao/tour/blog/global"
	"giao/tour/blog/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string
	AppSecret string
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	expireTime := time.Now().Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMd5(appKey),
		AppSecret: util.EncodeMd5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := withClaims.SignedString(GetJwtSecret())
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
