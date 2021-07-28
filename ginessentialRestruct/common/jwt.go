package common

import (
	"ProjectPractice/ginessentialRestruct/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect") //用于加密的字符串

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*  协议信息.内容主题内容.hash(协议信息+内容主题内容+a_secret_crect)
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcsImV4cCI6MTYyNzgwMjI0NSwiaWF0IjoxNjI3MTk3NDQ1LCJp
c3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.16ZKkvuyDLAj9zkIdTJ3FyckqQ3wy3ujb_JORNyjLUs"
*/

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
