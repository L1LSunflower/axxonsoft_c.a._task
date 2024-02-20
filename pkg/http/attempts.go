package http

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/cerrors"
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/logger"
)

const Id = "uuid"

func Attempts(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		logInstance := logger.Instance()
		body, err := getRequestBody(req)
		if err != nil {
			ErrorResponse(resp, err.Error(), http.StatusBadRequest)
			return
		}
		requestId := req.Header.Get(Id)
		BeforeAttempt(logInstance, requestId, body)
		next.ServeHTTP(resp, req)
	})
}

func BeforeAttempt(logInstance *logger.LInstance, requestId string, body []byte) {
	message := &logger.Message{
		Message:     "ATTEMPT BEFORE",
		FullMessage: string(body),
		Datetime:    time.Now().Unix(),
		RequestId:   requestId,
	}
	logInstance.Info(message)
}

func AfterAttempt(logInstance *logger.LInstance, requestId string, body []byte) {
	message := &logger.Message{
		Message:     "ATTEMPT AFTER",
		FullMessage: string(body),
		Datetime:    time.Now().Unix(),
		RequestId:   requestId,
	}
	logInstance.Info(message)
}

func getRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, cerrors.ErrReadBody
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func getResponseBody(resp http.ResponseWriter) (*bytes.Buffer, error) {
	body := new(bytes.Buffer)
	_ = io.MultiWriter(resp, body)
	return body, nil
}
