package internal

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func AuthorizeUser() {

}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenstr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}
