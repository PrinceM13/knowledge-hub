package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrinceM13/knowledge-hub-api/internal/config"
	"github.com/PrinceM13/knowledge-hub-api/internal/db"
	"github.com/PrinceM13/knowledge-hub-api/internal/db/user"
	"github.com/PrinceM13/knowledge-hub-api/internal/server"
)

func main() {
	cfg := config.Load()

	if err := db.Connect(cfg); err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	userRepo := user.NewPostgresRepository(db.DB)

	app := &server.App{
		UserRepo: userRepo,
	}

	engine := server.New(app)

	addr := ":" + cfg.Port
	log.Printf("🚀 API server running on port %s (env=%s)\n", addr, cfg.AppEnv)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// run server in a goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %s\n", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("error closing database: %v\n", err)
	}

	log.Println("✅ Server exited gracefully")
}
