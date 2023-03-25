package main

import (
	"concurrency-chat/Logger"
	"concurrency-chat/app"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		Logger.ErrorLogger().Printf("error loading .env file: %v", err)
		return
	}
	// handle multiple request
	// speaking with goroutines
	// what

	http.HandleFunc("/", app.HandleChat)
	go app.HandleMessage()

	go http.ListenAndServe(":8080", nil)

	select {}
}
