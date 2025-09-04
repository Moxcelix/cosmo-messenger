package pkg

import (
	"main/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt struct {
	secret string
}

func NewJwt(env config.Env) *Jwt {
	return &Jwt{
		secret : env.JwtSecreet
	}
}

func (j *Jwt) GenerateToken(userId string, minutes int32) (string, error){
	claims := jwt.MapClaims{
		"userID": userId,
		"exp": time.Now().Add(minutes * time.Minute).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *Jwt) ValidateToken (tokenStr string) (string, error){
	token, err:= jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error){
		return []byte(j.secret), nil
	})

	if err != nil || !token.Valid{
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["userID"].(string), nil
}
