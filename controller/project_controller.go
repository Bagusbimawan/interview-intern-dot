package controller

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/service"
	"intervew-intern-dot/utils"
	"net/http"
	"strconv"

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

func GetProjectByID(c *gin.Context) {
	idParam := c.Param("id") // Ambil ID dari URL
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	project, err := service.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func UpdateProjectByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedProject model.Project
	if err := c.ShouldBindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = service.UpdateProjectByID(id, &updatedProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project berhasil diperbarui", "data": updatedProject})
}
