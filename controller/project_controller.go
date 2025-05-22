package controller

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/service"
	"intervew-intern-dot/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	UserID := utils.GetUserIDFromContext(c)
	var input model.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	input.UserID = uint(UserID)

	// Simpan project ke database lewat service
	if err := service.CreateProject(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "project created successfully",
		"data":    input,
	})
}
