package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SecretKey 用于签名和验证 JWT 的密钥
var SecretKey = []byte("your-secret-key")

// validateToken 使用 JWT 验证 token 的有效性
func validateToken(tokenString string) bool {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是 HMAC 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretKey, nil
	})

	// 检查解析是否成功且 token 是否有效
	return err == nil && token.Valid
}

// AuthMiddleware 是一个简单的鉴权中间件示例
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 提取 token 部分
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// 验证 token 是否有效
		if !validateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 解析 token 并提取 claims
		claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 确保使用的是 HMAC 签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return SecretKey, nil
		})

		if err != nil || !claims.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 将 claims 存储在 gin.Context 中
		c.Set("claims", claims.Claims)

		// token 有效，继续处理请求
		c.Next()

	}
}
