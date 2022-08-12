package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/models/jwtModel"
	"time"
)

type JWTClaims struct {
	Info jwtModel.UserInfo
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 12

// GenToken 生成JWT
func GenToken(info jwtModel.UserInfo) (string, error) {
	c := JWTClaims{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    config.GetConfig().Auth.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(config.GetConfig().Auth.Secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.GetConfig().Auth.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
