package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)
	return token.SignedString([]byte(config.SecretKey))
}
