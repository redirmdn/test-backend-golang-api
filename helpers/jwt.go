package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	now := time.Now()
	expiresAt := now.Add(10 * time.Minute)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})

	token, err := tokenClaims.SignedString(secretKey)
	return token, err
}

func ParseJWT(token string) (*Claims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
