package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Gunakan secret hardcoded yang sama dengan yang digunakan saat login
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			fmt.Println("Token parse error:", err.Error())
			c.JSON(401, gin.H{"error": "invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(401, gin.H{"error": "token not valid"})
			c.Abort()
			return
		}

		// Ekstrak klaim dari token dan simpan ke context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			if userID, exists := claims["user_id"]; exists {

				c.Set("user_id", int(userID.(float64)))
			}
		}

		c.Next()
	}
}
