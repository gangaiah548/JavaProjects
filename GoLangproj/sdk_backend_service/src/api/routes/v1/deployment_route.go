package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func DeploymentRouter(timeout time.Duration, group *gin.RouterGroup) error {

	// use timeout if required at the controller level
	/*var deploymentControllerObject, err = controller.NewProcessDeploymentController()

	if err != nil {
		return err
	}
	group.POST("/deploy", deploymentControllerObject.Deployment)
	*/
	return nil
}
