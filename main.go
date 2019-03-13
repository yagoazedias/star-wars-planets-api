package main

import (
	"github.com/yagoazedias/star-wars-planets-api/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "4000"
	}

	server := service.NewServer()
	server.Run(":"  + port)
}

