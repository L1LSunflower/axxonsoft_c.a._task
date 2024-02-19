package entities

type Task struct {
	Id             string         `json:"id"`
	Status         string         `json:"status"`
	HttpStatusCode int            `json:"httpStatusCode"`
	Headers        map[string]any `json:"headers"`
	ContentLength  int            `json:"length"`
}
