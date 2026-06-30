package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/regis37/ssh-portfolio/internal/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatal("Failed to create server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting SSH portfolio", "addr", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, net.ErrClosed) {
			log.Fatal("Server error", "error", err)
		}
	}()

	<-done
	log.Info("Shutting down…")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Shutdown error", "error", err)
	}
}
