package main

import (
	"log"

	"github.com/PrinceM13/knowledge-hub-api/internal/server"
)

func main() {
	r := server.New()

	log.Println("🚀 API server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
