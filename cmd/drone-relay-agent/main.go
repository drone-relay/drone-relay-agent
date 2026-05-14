package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/drone-relay/drone-relay-agent/internal/collector"
	"github.com/drone-relay/drone-relay-agent/internal/config"
	"github.com/drone-relay/drone-relay-agent/internal/output"
)

func main() {
	cfg := config.Default()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	hostCollector, err := collector.NewHostCollector()
	if err != nil {
		log.Fatalf("failed to initialize host collector: %v", err)
	}

	jsonWriter := output.NewJSONWriter(os.Stdout)

	ticker := time.NewTicker(cfg.CollectionInterval)
	defer ticker.Stop()

	log.Println("drone-relay-agent started")

	for {
		select {
		case <-ctx.Done():
			log.Println("drone-relay-agent stopped")
			return

		case <-ticker.C:
			metrics, err := hostCollector.Collect(ctx)
			if err != nil {
				log.Printf("failed to collect host metrics: %v", err)
				continue
			}

			if err := jsonWriter.Write(metrics); err != nil {
				log.Printf("failed to write metrics: %v", err)
			}
		}
	}
}