package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.luizalabs.com/taste-match-api/internal/config"
)

// WaitingForShutdown register a channel and waiting for finalization signal
func WaitingForShutdown(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGKILL)
	log.Println(<-ch)

	log.Println("Starting server shutdown...")

	wait := time.Duration(config.Get().Server.TimeToCompleteShutdown) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server shutdown completed")
}
