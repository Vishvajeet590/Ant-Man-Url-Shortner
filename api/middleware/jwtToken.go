package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenCalims struct {
	Role string
	Id   int
	jwt.StandardClaims
}

var secretJwtKey = []byte("AllYourBase")

func GenerateToken(id int, roles string) (string, error) {

	claims := &TokenCalims{
		roles,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(4 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretJwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func AuthenticateToken(jwtToken string) (bool, int, string, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &TokenCalims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretJwtKey, nil
	})

	if err != nil {
		return false, -999, "", err
	}

	if claims, ok := token.Claims.(*TokenCalims); ok && token.Valid {
		return true, claims.Id, claims.Role, nil
	} else {
		return false, -999, "", err
	}
}
