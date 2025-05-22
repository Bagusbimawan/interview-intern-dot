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

		// Log header Authorization
		fmt.Println("Auth Header:", authHeader)

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Log token yang akan diparse
		fmt.Println("Token to parse:", tokenString)

		// Log JWT_SECRET (hanya untuk debugging, jangan gunakan di production)
		fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Log signing method
			fmt.Println("Token signing method:", token.Method.Alg())
			// Gunakan secret hardcoded yang sama dengan yang digunakan saat login
			return []byte("bagus"), nil
			// return []byte(os.Getenv("JWT_SECRET")), nil
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
			fmt.Println("Token claims:", claims)
			if userID, exists := claims["user_id"]; exists {
				fmt.Println("User ID from token:", userID)
				c.Set("user_id", int(userID.(float64)))
			} else {
				fmt.Println("user_id claim not found in token")
			}
		}

		c.Next()
	}
}
