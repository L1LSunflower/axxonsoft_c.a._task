package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/cerrors"
	httpHandlers "github.com/L1LSunflower/axxonsoft_c.a._task/pkg/http"
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/logger"
)

const Id = "uuid"

func Attempts(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		logInstance := logger.Instance()
		body, err := getRequestBody(req)
		if err != nil {
			httpHandlers.ErrorResponse(resp, err.Error(), http.StatusBadRequest)
			return
		}
		message := &logger.Message{
			Message:     "ATTEMPT BEFORE",
			FullMessage: string(body),
			Datetime:    time.Now().Unix(),
			ExtraData:   nil,
		}
		logInstance.Info(message)
		req.Header.Add(Id, uuid.New().String())
		next.ServeHTTP(resp, req)
		body, err = getResponseBody(resp)
		if err != nil {
			httpHandlers.ErrorResponse(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		message = &logger.Message{
			Message:     "ATTEMPT AFTER",
			FullMessage: string(body),
			Datetime:    time.Now().Unix(),
			ExtraData:   nil,
		}
		logInstance.Info(message)
	})
}

func getRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, cerrors.ErrReadBody
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func getResponseBody(resp http.ResponseWriter) ([]byte, error) {
	var body bytes.Buffer
	_ = io.MultiWriter(resp, &body)
	return body.Bytes(), nil
}
