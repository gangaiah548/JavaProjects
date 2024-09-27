package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ProcessLifecycleRouter(timeout time.Duration, group *gin.RouterGroup) {

	// use timeout if required at the controller level

	//group.POST("/startProcess", controller.NewProcessLifecycleController().StartProcess)
	//group.POST("/publishMsg", controller.NewProcessLifecycleController().Event)
}
