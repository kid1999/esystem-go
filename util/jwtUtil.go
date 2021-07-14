/**
@auther: kid1999
@date: 2021/7/14 15:02
@desc: jwtUtil
**/

package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//var jwtSecret = []byte(config.Get().JwtSecret)

var jwtSecret = []byte("dasd65456dasdx")

type Claims struct {
	StudentID string `json:"student_id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// create a new token
func GenerateToken(student_id, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		student_id,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解码和校验
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


