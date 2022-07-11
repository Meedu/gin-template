package types

import (
	"github.com/golang-jwt/jwt/v4"
)

type LaravelJWTPayload struct {
	Subject int64  `json:"sub"`
	Prv     string `json:"prv"`
	jwt.StandardClaims
}
