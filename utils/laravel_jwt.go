package utils

import (
	"errors"
	"time"

	"github.com/Meedu/gin-template/global"
	"github.com/Meedu/gin-template/types"
	"github.com/golang-jwt/jwt/v4"
)

type LaravelJWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *LaravelJWT {
	return &LaravelJWT{
		SigningKey: []byte(global.MD_CONFIG.JWT.SigningKey),
	}
}

func (j *LaravelJWT) NewPayload(userId int64, prv string) types.LaravelJWTPayload {
	claims := types.LaravelJWTPayload{
		Subject: userId,
		Prv:     prv, //签发jwt的Model
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.MD_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.MD_CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

func (j *LaravelJWT) NewToken(claims types.LaravelJWTPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *LaravelJWT) ParseToken(tokenString string) (*types.LaravelJWTPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &types.LaravelJWTPayload{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*types.LaravelJWTPayload); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, TokenInvalid
}
