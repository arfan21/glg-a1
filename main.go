package main

import (
	"github.com/arfan21/hacktiv8-golang-jwt/router"
)

func main() {
	server := router.StartApp()

	server.Run(":9000")
}
