package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"e1rms/internal/e1rm/handler"
	"e1rms/internal/e1rm/service"
)

func main() {
	e1rmService := e1rm_service.New()
	e1rmHandler := e1rm_handler.New(e1rmService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/e1rm", e1rmHandler.ServeE1rmRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listing on port %s", port)
	log.Printf("Error? %s", http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
