package main

import (
	"go-web-template/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(cfg.ServerConfig.Port, nil)
}
