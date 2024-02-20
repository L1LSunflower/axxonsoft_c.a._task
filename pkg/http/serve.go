package http

import (
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Serve(port int) error {
	router := mux.NewRouter()
	router.Handle("/task", RequestId(Attempts(http.HandlerFunc(RegisterTask))))
	router.Handle("/task/{"+Id+"}", RequestId(Attempts(http.HandlerFunc(Task))))
	strPort := strconv.Itoa(port)
	logger.Instance().Info(&logger.Message{Message: "SERVER STARTED ON PORT: " + strPort})
	if err := http.ListenAndServe(":"+strPort, router); err != nil {
		return err
	}
	return nil
}
