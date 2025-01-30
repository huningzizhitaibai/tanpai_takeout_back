package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"tanpai_takeout_back/common/enum"
	"time"
)

//定义jwt生成中需要包含的参数

type CustomClaims struct {
	Username string        `json:"username"`
	Type     enum.UserType `json:"type"` //用户权限
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(username string, _type enum.UserType, secret string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Type:     _type,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "huning",
			Subject:   enum.IntT2StrT(_type),
			Audience:  jwt.ClaimStrings{"PC", "Wechat_Program"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 50)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseToken(tokenString string, secret string) (*CustomClaims, error) {
	parseToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parseToken.Claims.(*CustomClaims); ok && parseToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
