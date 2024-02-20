package entities

type RegisterTask struct {
	Url     string           `json:"url"`
	Method  string           `json:"method"`
	Headers []map[string]any `json:"headers"`
}

func (t *RegisterTask) ToTask(id string) *Task {
	return &Task{
		Id:      id,
		Url:     t.Url,
		Method:  t.Method,
		Headers: t.Headers,
	}
}
