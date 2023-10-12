package entity

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	UserID   string `json:"user_id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	jwt.StandardClaims
}
