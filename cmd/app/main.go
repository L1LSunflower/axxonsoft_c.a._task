package main

import (
	"github.com/L1LSunflower/axxonsoft_c.a._task/config"
	"github.com/L1LSunflower/axxonsoft_c.a._task/pkg/http"
	"log"
)

func main() {
	cfg := config.GetConfig()
	if err := http.Serve(cfg.Port); err != nil {
		log.Println("failed to start server" + err.Error())
	}
}
