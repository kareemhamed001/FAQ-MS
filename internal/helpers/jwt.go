package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken signs a JWT with provided claims plus sensible defaults.
// Adds iat/exp if they are missing; default TTL is 24h.
func GenerateToken(data map[string]interface{}, secret string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{}

	for key, value := range data {
		claims[key] = value
	}

	if _, ok := claims["iat"]; !ok {
		claims["iat"] = now.Unix()
	}

	if _, ok := claims["exp"]; !ok {
		claims["exp"] = now.Add(24 * time.Hour).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken parses and validates a JWT (signature + registered claims).
func ValidateToken(tokenString string, secret string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
