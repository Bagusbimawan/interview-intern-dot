package utils

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) int {
	UserID, exist := c.Get("user_id")
	if !exist {
		return 0
	}

	return UserID.(int)
}
