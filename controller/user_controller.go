package controller

import (
	"intervew-intern-dot/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	user, err := service.GetUserProfile(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user found", "data": gin.H{
		"email": user.Email,
		"name":  user.Name,
	}})
}
