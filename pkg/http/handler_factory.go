package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/L1LSunflower/axxonsoft_c.a._task/internal/entities"
)

const (
	contentTypeKey = "content-type"
	contentTypeVal = "application/json"
)

func ErrorResponse(resp http.ResponseWriter, errorMessage string, statusCode int) {
	errResp, err := entities.NewErrorResponse(errorMessage).ToBytes()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Write(errResp)
	resp.Header().Set(contentTypeKey, contentTypeVal)
	resp.WriteHeader(statusCode)
}

func RegisterTask(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		ErrorResponse(resp, "that method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if cTypeVal := req.Header.Get(contentTypeKey); cTypeVal != contentTypeVal {
		ErrorResponse(resp, "wrong content type", http.StatusBadRequest)
		return
	}
}

func Task(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		ErrorResponse(resp, "that method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	id, ok := vars["taskId"]
	if !ok {
		ErrorResponse(resp, "task is missing in parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(id)
}
