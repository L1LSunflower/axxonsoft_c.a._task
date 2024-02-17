package logger

type Message struct {
	Message     string
	FullMessage string
	Datetime    int
	ExtraData   map[string]any
}
