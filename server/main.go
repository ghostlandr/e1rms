package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	e1rm_handler "e1rms/internal/e1rm/handler"
	e1rm_model "e1rms/internal/e1rm/model"
	e1rm_service "e1rms/internal/e1rm/service"
)

func main() {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	e1rmModel := e1rm_model.New(conn)
	e1rmService := e1rm_service.New(e1rmModel)
	e1rmHandler := e1rm_handler.New(e1rmService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/e1rm", e1rmHandler.ServeE1rmRequest)
	mux.HandleFunc("/api/v0/e1rms", e1rmHandler.ServeListE1rmRequest)
	mux.HandleFunc("/api/v0/provision_e1rm_table", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		e1rmModel.ProvisionTables(ctx)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listing on port %s", port)
	log.Printf("Error? %s", http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
