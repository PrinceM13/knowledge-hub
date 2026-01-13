package main

import (
	"log"
	"net/http"

	"github.com/PrinceM13/knowledge-hub-api/internal/health"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", health.Handler)

	server := &http.Server{Addr: ":8080", Handler: mux}

	log.Println("🚀 API server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
