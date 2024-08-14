package dtos

import "sdk_backend_service/src/models"

type ProductResponseDto struct {
	StatusCode int                           `json:"statusCode"`
	Message    string                        `json:"message"`
	Data       map[string]interface{}        `json:"data"`
	Deployment models.ProcessDeploymentModel `json:"deployment"`
}
