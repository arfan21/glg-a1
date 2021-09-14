package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arfan21/hacktiv8-assignment-2/config"
)

func Start() {
	cfg := config.New()
	// mysqlClient := mysql.New(cfg)
	// defer mysqlClient.Close()

	port := cfg.Get("PORT")
	if port == "" {
		port = "8000"
	}
	baseUrl := cfg.Get("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost"
	}

	muxRouter := NewRouter()
	log.Println("HTTP server listening on " + baseUrl + ":" + port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), muxRouter)
}
