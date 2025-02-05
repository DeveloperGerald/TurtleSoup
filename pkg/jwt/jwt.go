package jwt

import (
	"fmt"
	"time"

	"github.com/DeveloperGerald/TurtleSoup/model"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT 密钥，用于签名和验证
var JwtKey = []byte("turtle_soup_b8r5s4c5d5b5re5f5d")

// UserClaims 用于定义 JWT 中的自定义字段
type UserClaims struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	jwt.StandardClaims
}

// GenerateJWT 用于生成一个新的 JWT
func GenerateJWT(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置 JWT 过期时间为 24 小时

	claims := &UserClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用密钥对 JWT 进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("jwt signed error: %v", err)
	}

	return signedToken, nil
}
