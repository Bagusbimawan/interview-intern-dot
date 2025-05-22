package controller

import (
	"intervew-intern-dot/model"
	"intervew-intern-dot/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller/task_controller.go
func CreateTaskUnderProject(c *gin.Context) {
	projectIDParam := c.Param("id")
	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid project id"})
		return
	}

	// Validasi apakah project exist
	project, err := service.GetProjectDetailByID(projectID)
	if err != nil || project == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	task.ProjectID = uint(projectID)

	if err := service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task created successfully", "data": task})
}
