package service

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

type Claims struct {
	UID uint64
	jwt.RegisteredClaims
}

func GenerateToken(userID uint64) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	now := time.Now()
	expire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	expireTime := now.Add(time.Second * time.Duration(expire))
	claims := Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func ParseToken(tokenString string) (*Claims, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*Claims)
	return claims, nil
}

func GetUID(tokenString string) (uint64, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UID, nil
}
