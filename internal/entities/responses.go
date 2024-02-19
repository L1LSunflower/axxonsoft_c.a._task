package entities

import "encoding/json"

const errorStatus = "error"

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func (e *ErrorResponse) ToBytes() ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewErrorResponse(errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		Status:       errorStatus,
		ErrorMessage: errorMessage,
	}
}
