package main

import (
	"e1rms/internal/e1rm"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/e1rm", e1rm.NewHandler().ServeE1rmRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listing on port %s", port)
	log.Printf("Error? %s", http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
