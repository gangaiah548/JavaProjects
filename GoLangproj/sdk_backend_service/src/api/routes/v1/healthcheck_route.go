package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(timeout time.Duration, group *gin.RouterGroup) {

	// use timeout if required at the controller level

	//group.GET("/healthCheck", controller.NewHealthCheckController().HealthCheck)
}
