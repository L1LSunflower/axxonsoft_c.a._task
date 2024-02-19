package entities

type RegisterTask struct {
	Url     string         `json:"url"`
	Method  string         `json:"method"`
	Headers map[string]any `json:"headers"`
}
