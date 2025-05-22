package routes

import (
	"intervew-intern-dot/controller"
	"intervew-intern-dot/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/user/:id", controller.GetUserProfile)

	auth := r.Group("/api")

	auth.Use(middleware.JWTAuthMiddleware())
	auth.POST("/project", controller.CreateProject)
	auth.GET("/project/:id", controller.GetProjectByID)
	auth.POST("/project/:id/task", controller.CreateTaskUnderProject)
	auth.PUT("/project/:id", controller.UpdateProjectByID)
}
