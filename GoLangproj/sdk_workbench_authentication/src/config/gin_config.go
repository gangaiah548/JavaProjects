package config

import (
	"github.com/gin-gonic/gin"

	"sdk_workbench_authentication/src/api/middleware"
)

func SetGinMode(ginMode string) {
	// Initialize Gin mode
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func InitGin(env string) *gin.Engine {
	r := gin.New()                                 // empty engine
	r.Use(middleware.DefaultStructuredLogger(env)) // adds  new middleware
	r.Use(middleware.Recovery)                     // adds the default recovery middleware

	return r
}
