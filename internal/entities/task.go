package entities

import (
	"encoding/json"
)

type Task struct {
	Id             string           `json:"id"`
	Url            string           `json:"url"`
	Method         string           `json:"method"`
	Status         Status           `json:"status"`
	HttpStatusCode int              `json:"httpStatusCode"`
	Headers        []map[string]any `json:"headers"`
	ContentLength  int              `json:"length"`
}

func (t *Task) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

type Status string

const (
	NewStatus     Status = "new"
	ErrorStatus   Status = "error"
	ProcessStatus Status = "in_process"
	DoneStatus    Status = "done"
)

var StatusesOrder = []Status{NewStatus, ErrorStatus, ProcessStatus, DoneStatus}

func DefineStatus(status string) Status {
	switch Status(status) {
	case NewStatus:
		return NewStatus
	case ErrorStatus:
		return ErrorStatus
	case ProcessStatus:
		return ProcessStatus
	case DoneStatus:
		return DoneStatus
	default:
		return ErrorStatus
	}
}
