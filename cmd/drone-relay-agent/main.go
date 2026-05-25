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
    	log.Fatalf("unsupported output format: %s", cfg.OutputFormat)
    }

    if cfg.BatchSize <= 0 {
    	log.Fatalf("batch size must be greater than 0")
    }

    if cfg.CollectionInterval <= 0 {
    	log.Fatalf("collection interval must be greater than 0")
    }

    if cfg.FlushInterval <= 0 {
    	log.Fatalf("flush interval must be greater than 0")
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

	collectionTicker := time.NewTicker(cfg.CollectionInterval)
    defer collectionTicker.Stop()

    flushTicker := time.NewTicker(cfg.FlushInterval)
    defer flushTicker.Stop()

    batch := make([]model.HostMetrics, 0, cfg.BatchSize)

	log.Println("drone-relay-agent started")

    for {
        select {
        case <-ctx.Done():
           flushBatch(jsonWriter, batch, "shutdown")
            log.Println("drone-relay agent stopped")
            return

        case <-collectionTicker.C:
            metrics, err := hostCollector.Collect(ctx)
            if err != nil {
                log.Printf("failed to collect host metrics: %v", err)
                continue
            }

            batch = append(batch, metrics)

            if len(batch) >= cfg.BatchSize {
                flushBatch(jsonWriter, batch, "batch-size")
                batch = batch[:0]
            }

        case <-flushTicker.C:
            if len(batch) > 0 {
                flushBatch(jsonWriter, batch, "flush-interval")
                batch = batch[:0]
            }
        }
    }
}

func flushBatch(
    writer *output.JSONWriter,
    batch []model.HostMetrics,
    reason string,
) {
    if len(batch) == 0 {
        return
    }

    log.Printf("flushing metrics batch size=%d reason=%s", len(batch), reason)

    if err := writer.WriteBatch(batch); err != nil {
        log.Printf("failed to write metrics batch: %v", err)
    }
}
