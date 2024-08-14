package dtos

import "sdk_workbench_authentication/src/models"

type DeploymentResponseDto struct {
	StatusCode int                           `json:"statusCode"`
	Message    string                        `json:"message"`
	Data       map[string]interface{}        `json:"data"`
	Deployment models.ProcessDeploymentModel `json:"deployment"`
}
