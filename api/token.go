package api

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SigningKey = "serdfgyhjkkljhvgh"

type tokenClaims struct {
	jwt.StandardClaims
	Email string
}

func (s *Server) GenerateToken(email, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 + time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		email,
	})

	return token.SignedString([]byte(SigningKey))
}

func (s *Server) VerifyToken(accessToken string) (string, error) {
	claims := &tokenClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return SigningKey, nil
	})

	if token != nil || !token.Valid {
		return claims.Email, nil
	}

	return "", err
}
