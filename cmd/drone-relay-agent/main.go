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
	"github.com/drone-relay/drone-relay-agent/internal/model"
)

func main() {
	cfg := config.Load()

    if cfg.OutputFormat != "json" {
        log.Fatalf(
            "unsupported output format: %s",
            cfg.OutputFormat,
        )
    }

    if cfg.BatchSize <= 0 {
        log.Fatalf("batch size must be greater than 0")
    }

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

    log.Printf(
        "starting drone-relay-agent interval=%s output=%s",
        cfg.CollectionInterval,
        cfg.OutputFormat,
    )

	ticker := time.NewTicker(cfg.CollectionInterval)
	defer ticker.Stop()

	log.Println("drone-relay-agent started")

    batch := make([]model.HostMetrics, 0, cfg.BatchSize)

	for {
		select {
		case <-ctx.Done():
            if len(batch) > 0 {
                    if err := jsonWriter.WriteBatch(batch); err != nil {
                        log.Printf("failed to flush metrics batch during shutdown: %v", err)
                    }
                }

                log.Println("drone-relay-agent stopped")
                return

        case <-ticker.C:
            metrics, err := hostCollector.Collect(ctx)
            if err != nil {
                log.Printf("failed to collect host metrics: %v", err)
                continue
            }

            batch = append(batch, metrics)

            if len(batch) >= cfg.BatchSize {
                if err := jsonWriter.WriteBatch(batch); err != nil {
                    log.Printf("failed to write metrics batch: %v", err)
                    continue
                }

                batch = batch[:0]
            }
		}
	}
}