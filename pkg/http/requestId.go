package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		requestId, ok := vars[Id]
		if !ok {
			requestId = uuid.New().String()
			req.Header.Add(Id, requestId)
		} else {
			req.Header.Add(Id, requestId)
		}
		next.ServeHTTP(resp, req)
	})
}
