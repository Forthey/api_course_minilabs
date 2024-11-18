package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/logger"
	"server/routers"
)

func createServer() {
	mux := http.NewServeMux()

	routers.Init(mux)

	loggingMux := logger.NewLogger(mux)

	fmt.Printf("Starting server on port %s\n", config.Settings.ListenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Settings.ListenPort), loggingMux)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	createServer()
}
