package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(userId int) (string, error)  {
	claims := jwt.MapClaims{
		"userid": int64(userId),
		"role": "admin",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	//JWT := os.Getenv("JWT_SECRET")
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := tokenWithClaims.SignedString([]byte("JWT"))
	return tk, err
}
