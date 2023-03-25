package gjwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

func GetJwtToken(uid string, role string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UID:  uid,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(SECRET))
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", errors.New("Token 过期")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", "", errors.New("Error: Invalid claims")
	}
	return claims.UID, claims.Role, nil
}
