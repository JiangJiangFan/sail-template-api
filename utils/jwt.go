package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func JwtToken(obj map[string]interface{}) (string,error) {
	// 生成JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(obj))

	// 设置Token的Claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = u.Username
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 生成Token字符串
	tokenString, err := token.SignedString([]byte("secret"))
	return tokenString, err
}

func ValidateJwtToken(tokenString string) map[string]interface{} {
	// 解析JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil

	})
	if err != nil {
		return nil
	}

	// 验证Token是否有效
	claims := token.Claims.(jwt.MapClaims)
	return claims
}
