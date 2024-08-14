package controller

/*import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sdk_backend_service/src/config/logger"
	"sdk_backend_service/src/dtos"
)

type IHealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type HealthCheckControllerImpl struct {
}

func NewHealthCheckController() IHealthCheckController {
	return &HealthCheckControllerImpl{}
}

func (hcc *HealthCheckControllerImpl) HealthCheck(c *gin.Context) {

	// put the health check logic for the application

	logger.Info().Msg("Controller invoked properly")

	hc := dtos.HealthCheckDto{
		HttpStatus:        http.StatusOK,
		ApplicationStatus: "OK",
	}

	c.JSON(http.StatusOK, hc)
}
*/
