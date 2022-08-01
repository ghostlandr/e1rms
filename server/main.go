package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	e1rm_handler "e1rms/internal/e1rm/handler"
	e1rm_model "e1rms/internal/e1rm/model"
	e1rm_service "e1rms/internal/e1rm/service"
)

func main() {
	dburl := os.Getenv("DATABASE_URL")
	if dburl != "" {
		fmt.Println("Got a DB url that was longer than an empty string...")
	}
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("We seem to have gotten through the DB connection phase")

	e1rmModel := e1rm_model.New(conn)
	e1rmService := e1rm_service.New(e1rmModel)
	e1rmHandler := e1rm_handler.New(e1rmService)

	fmt.Println("Domain objects created")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/e1rm", e1rmHandler.ServeE1rmRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listing on port %s", port)
	log.Printf("Error? %s", http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
