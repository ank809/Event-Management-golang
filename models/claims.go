package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}
