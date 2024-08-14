package controller

/*import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/dtos"
	"sdk_backend_service/src/services"
)

type IProcessDeploymentController interface {
	Deployment(c *gin.Context)
}

type ProcessDeploymentControllerImpl struct {
	processDeploymentService services.IProcessDeploymentService
}

func NewProcessDeploymentController() (IProcessDeploymentController, error) {
	var processDeploymentService, err = services.NewProcessDeploymentService()

	if err != nil {
		return nil, err
	}
	return &ProcessDeploymentControllerImpl{
		processDeploymentService: processDeploymentService,
	}, nil
}

func (pdc *ProcessDeploymentControllerImpl) Deployment(c *gin.Context) {
	logger.Info().Msg("Deployment service invoked properly")
	//upload
	processDeploymentModel, err := pdc.processDeploymentService.PrepareDeploymentModel(c.Request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.DeploymentResponseDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       map[string]interface{}{"data": err.Error()},
			})
		return
	}

	deploymentResponse, err := pdc.processDeploymentService.CreateProcessDeployment(processDeploymentModel)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.DeploymentResponseDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       map[string]interface{}{"data": err.Error()},
			})
		return
	}

	// returning empty string since process definition as return is not required as it will unnecessarily increase the return payload size
	deploymentResponse.Definition = ""

	c.JSON(
		http.StatusOK,
		dtos.DeploymentResponseDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       map[string]interface{}{"data": "success"},
			Deployment: deploymentResponse,
		})

}*/

//TODO Write a deployment updater
