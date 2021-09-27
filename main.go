package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arfan21/hacktiv8-assignment-3/app"
)

func main() {
	PORT := ":8000"
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	go func() {
		_ = app.Reloader(app.ChangeWeatherStatus, 5*time.Second)
	}()

	fmt.Println("HTTP server listening on http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
