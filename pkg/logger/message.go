package logger

import "encoding/json"

type Message struct {
	Message     string `json:"message"`
	FullMessage string `json:"full-message,omitempty"`
	Datetime    int64  `json:"datetime,omitempty"`
	RequestId   string `json:"request-id,omitempty"`
}

func (m *Message) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
