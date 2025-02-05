package middleware

import (
	"net/http"
	"strings"

	"github.com/DeveloperGerald/TurtleSoup/consts"
	"github.com/DeveloperGerald/TurtleSoup/controller"
	myjwt "github.com/DeveloperGerald/TurtleSoup/pkg/jwt"
	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)

// Middleware 用于验证 JWT 的中间件
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, controller.StandardResp{Code: -1, Message: "Unauthorized"})
			c.Abort()
			return
		}

		// 提取 Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, controller.StandardResp{Code: -1, Message: "Unauthorized"})
			c.Abort()
			return
		}

		// 解析 Token
		claims := &myjwt.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// 使用密钥验证 Token 的签名
			return myjwt.JwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, controller.StandardResp{Code: -1, Message: "Authorize failed"})
			c.Abort()
			return
		}

		c.Set(consts.UserClaimsKey, claims)
		c.Next()
	}
}
