package common

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("POSKWDJAWNDMSNawdaw")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.UserInfo) (string, error) {
	exptime := time.Now().Add(7 * 4 * time.Hour)

	claims := &Claims{
		UserID: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exptime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "TGPServer",
			Subject:   "UserInfo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger := logging.GetLog()
		logger.Error("JWT组件新建Token失败")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
