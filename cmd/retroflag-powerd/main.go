package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("retroflag-powerd starting")
	log.Println("retroflag-powerd ready")

	<-ctx.Done()

	log.Println("shutdown signal received")
	log.Println("retroflag-powerd stopped")
}
