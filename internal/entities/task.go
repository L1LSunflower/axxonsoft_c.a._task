package entities

type Task struct {
	Id             string    `json:"id"`
	Status         string    `json:"status"`
	HttpStatusCode int       `json:"httpStatusCode"`
	Headers        []*Header `json:"headers"`
	ContentLength  int       `json:"length"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
