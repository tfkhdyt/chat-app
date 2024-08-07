package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	PrivateKey   string `json:"private_key"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	ID uint `json:"id"`
}

func (j *JwtClaims) SetExp(t time.Time) {
	j.ExpiresAt = jwt.NewNumericDate(t)
	j.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	j.NotBefore = jwt.NewNumericDate(time.Now().UTC())
}
