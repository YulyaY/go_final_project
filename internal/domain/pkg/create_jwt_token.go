package pkg

import (
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(jwtT, secret string) error {
	jwtToken, err := jwt.Parse(jwtT, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !jwtToken.Valid {
		return err
	}

	return nil
}
