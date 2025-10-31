package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"-"`
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
