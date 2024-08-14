package dtos

type EventDto struct {
	Message           string                 `json:"message" validate:"required"`
	ProcessInstanceId int64                  `json:"processInstanceId" validate:"required"`
	EngineName        string                 `json:"engineName" validate:"required"`
	MsgData           map[string]interface{} `json:"msgData"`
}

type EventResponseDto struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
