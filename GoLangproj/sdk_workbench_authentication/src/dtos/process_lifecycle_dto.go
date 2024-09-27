package dtos

type StartProcessRequestDto struct {
	Key         string                 `json:"_key,omitempty" validate:"required"`
	ProcessData map[string]interface{} `json:"processData" validate:"-"`
	ExecMode    string                 `json:"execMode" validate:"required"`
}

type StartProcessResponseDto struct {
	ProcessInstanceId string                 `json:"processInstanceId,omitempty"`
	Status            string                 `json:"status"`
	EngineName        string                 `json:"engineName"`
	Data              map[string]interface{} `json:"data"`
}
