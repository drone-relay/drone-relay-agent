package config

import (
	"flag"
	"time"
)

type Config struct {
	CollectionInterval time.Duration
	OutputFormat       string
	BatchSize          int
}

func Load() Config {
	interval := flag.Duration(
    		"interval",
    		5*time.Second,
    		"metrics collection interval",
    	)

    	output := flag.String(
    		"output",
    		"json",
    		"output format (json)",
    	)

        batchSize := flag.Int(
            "batch-size",
            5,
            "number of metrics per batch",
        )

    	flag.Parse()

    	return Config{
    		CollectionInterval: *interval,
    		OutputFormat:       *output,
    		BatchSize:          *batchSize,
    	}
}