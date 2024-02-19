package logger

type Message struct {
	Message     string
	FullMessage string
	Datetime    int64
	ExtraData   map[string]any
}
