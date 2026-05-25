package config

import (
	"flag"
	"time"
)

type Config struct {
	CollectionInterval time.Duration
	OutputFormat       string
	BatchSize          int
	FlushInterval      time.Duration
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

       flushInterval := flag.Duration(
            "flush-interval",
            10*time.Second,
            "maximum time between batch flushes",
        )


    	flag.Parse()

    	return Config{
    		CollectionInterval: *interval,
    		OutputFormat:       *output,
    		BatchSize:          *batchSize,
    		FlushInterval:      *flushInterval,
    	}
}