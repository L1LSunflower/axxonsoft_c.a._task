package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/http/middlewares"
)

func Serve(port int) error {
	router := mux.NewRouter()
	router.Handle("/task", middlewares.Attempts(http.HandlerFunc(RegisterTask)))
	router.Handle("/task/{taskId}", middlewares.Attempts(http.HandlerFunc(Task)))
	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		return err
	}
	return nil
}
