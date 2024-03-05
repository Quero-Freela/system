package services

import (
	"github.com/Quero-Freela/system/server/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(user *models.User) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["aud"] = "https://querofrela.app.br"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["jti"] = user.ID
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "https://auth.querofrela.app.br"
	claims["nbf"] = time.Now().Unix()
	claims["sub"] = user.Email
	claims["name"] = user.Name
	claims["email"] = user.Email

	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}

	claims["roles"] = roles

	tokenString, err := token.SignedString(user.JwtSecret)

	return &tokenString, err
}
