package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	UserID uint `json:userId`
	jwt.StandardClaims
}

var jwtSecret = []byte(viper.GetString("app.jwtSecret"))

// 生成token的函数
func GenerateToken(userId uint) (string, error) {
	fmt.Println("jwtSecret", viper.GetString("app.jwtSecret"))
	jwtExpireTime := viper.GetInt("app.jwtExpireTime")
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtExpireTime) * time.Minute)
	claims := Claims{
		userId, //自行添加的信息
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "admin",
		},
	}
	//	生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		claims, _ := tokenClaims.Claims.(*Claims)
		return claims, nil
	}
	return nil, err
}
