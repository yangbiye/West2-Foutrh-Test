package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims1 struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type Claims2 struct {
	ManagerID uint `json:"manager_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken1(userID uint) (string, error) {
	nowTime := time.Now()
	exp := nowTime.Add(3 * time.Hour)

	claims := Claims1{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    "Video",
		},
	}

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims)
	token, err := tokenClaims.SignedString([]byte("Video"))

	return token, err
}

// ParseToken 解析Token
func ParseToken1(token string) (*Claims1, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims1{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Video"), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims1)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GenerateToken 生成Token
func GenerateToken2(managerID uint) (string, error) {
	nowTime := time.Now()
	exp := nowTime.Add(3 * time.Hour)

	claims := Claims2{
		managerID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    "Video",
		},
	}

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims)
	token, err := tokenClaims.SignedString([]byte("Video"))

	return token, err
}

// ParseToken 解析Token
func ParseToken2(token string) (*Claims2, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims2{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Video"), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims2)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
