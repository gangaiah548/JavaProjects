package controller

/*import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/dtos"
	"sdk_backend_service/src/services"
)

type IProcessLifecycleController interface {
	StartProcess(c *gin.Context)
	Event(c *gin.Context)
}

type ProcessLifecycleControllerImpl struct {
	processLifecycleService services.IProcessLifecycleService
}

func NewProcessLifecycleController() IProcessLifecycleController {
	return &ProcessLifecycleControllerImpl{
		processLifecycleService: services.NewProcessLifecycleService(),
	}
}

func (plc *ProcessLifecycleControllerImpl) StartProcess(c *gin.Context) {
	logger.Info().Msg("StartProcess controller invoked properly")
	//upload
	var reqBody dtos.StartProcessRequestDto
	c.Bind(&reqBody)
	logger.Debug().Msg("Bind Successful [" + reqBody.Key + "]")
	startProcessResponse, err := plc.processLifecycleService.StartProcess(reqBody.Key, reqBody.ProcessData, reqBody.ExecMode)
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

	logger.Info().Msg(fmt.Sprint(startProcessResponse))

	c.JSON(http.StatusOK, startProcessResponse)

}

func (plc *ProcessLifecycleControllerImpl) Event(c *gin.Context) {

	// put the Event logic for the application
	println("Event Controller invoked")
	logger.Info().Msg("Event controller invoked properly")
	//upload
	var reqBody dtos.EventDto
	c.Bind(&reqBody)
	logger.Debug().Msg("Bind Successful with Message [" + reqBody.Message + "]")
	startProcessResponse, err := plc.processLifecycleService.Event(reqBody.Message, reqBody.MsgData, reqBody.ProcessInstanceId, reqBody.EngineName)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.EventResponseDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
			})
		return
	}

	logger.Info().Msg(fmt.Sprint(startProcessResponse))

	c.JSON(http.StatusOK, startProcessResponse)
}
*/
