package model

import (
	"workbench/graphql-app/graph/resolver/scalar"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token     string      `json:"token"`
	ExpiredAt scalar.Date `json:"-"`
}

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
