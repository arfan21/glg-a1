package main

import (
	"github.com/arfan21/hacktiv8-golang-jwt/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := router.StartApp()

	server.Run(":9000")
}
