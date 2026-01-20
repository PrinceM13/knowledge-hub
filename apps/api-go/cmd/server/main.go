package main

import (
	"log"

	"github.com/PrinceM13/knowledge-hub-api/internal/config"
	"github.com/PrinceM13/knowledge-hub-api/internal/server"
)

func main() {
	cfg := config.Load()

	r := server.New()

	addr := ":" + cfg.Port
	log.Printf("🚀 API server running on port %s (env=%s)\n", addr, cfg.AppEnv)

	if err := r.Run(addr); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
