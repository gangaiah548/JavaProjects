package dtos

type RestTaskRequestDto struct {
	RequestType string      `json:"reqType"`
	RequestData interface{} `json:"requestData"`
}

type RestTaskResponseDto struct {
	ResponseId        string            `json:"resId"`
	ProcessDataString map[string]string `json:"processDataString"`
	ProcessDataInt    map[string]int64  `json:"ProcessDataInt"`
	ResponseData      interface{}       `json:"responseData"`
}
