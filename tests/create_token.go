package tests

import (
	"log"

	"github.com/golang-jwt/jwt"
)

func createJwtToken(secret string) string {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln("err jwtToken transform: ", err)
		return ""
	}
	return token
}
