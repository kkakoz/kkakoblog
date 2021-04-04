package cryption

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"kkako-blog/pkg/errcode"
	"time"
)

type MyClaims struct {
	UserId int `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("hsdioahsodjisodxz")

const TokenExpireDuration = time.Hour * 24 * 3

func GenToken(userId int, username string) (string, error) {
	c := MyClaims{
		UserId: userId, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "zhangsan", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", errcode.AuthTokenGenerateErr
	}
	return signedString, nil
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return viper.GetString("jwt.secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}