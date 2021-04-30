package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type UserClaims struct {
	Uid uint
	jwt.StandardClaims
}

//签发token
func ReleaseToken(uid uint) (string, error) {
	claims := &UserClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Subject:   "auth token",
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(viper.GetInt("token.life"))).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("token.secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(viper.GetString("token.secret")), nil
	})
	return token, claims, err
}
