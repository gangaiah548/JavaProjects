package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

/*func Setup(timeout time.Duration, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")

	// Public APIs
	HealthCheckRouter(timeout, publicRouterV1)
	DeploymentRouter(timeout, publicRouterV1)
	ProcessLifecycleRouter(timeout, publicRouterV1)

	// Protected APIs

	// Private APIs
}*/

func Setup(timeout time.Duration, routerSDK1 *gin.RouterGroup) {
	publicRouterV1 := routerSDK1.Group("") //segregate group here
	//router := gin.Default()
	// Public APIs
	//HealthCheckRouter(timeout, publicRouterV1)
	//DeploymentRouter(timeout, publicRouterV1)
	//ProcessLifecycleRouter(timeout, publicRouterV1)

	//ProductRouter(timeout, publicRouterV1)
	SmpdataRouter(timeout, publicRouterV1)
	// Protected APIs

	// Private APIs
}
