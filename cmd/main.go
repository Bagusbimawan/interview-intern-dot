package main

import (
	"intervew-intern-dot/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	config.Initdb()
	r := gin.Default()

	r.Run(":8080")
}
