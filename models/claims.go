package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Name           string
	StandardClaims jwt.StandardClaims
}

// Valid implements jwt.Claims.
func (c *Claims) Valid() error {
	panic("unimplemented")
}
