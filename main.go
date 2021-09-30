package main

import (
	"fmt"
	"os"

	"github.com/arfan21/hacktiv8-golang-jwt/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := router.StartApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server.Run(fmt.Sprintf(":%s", port))
}
