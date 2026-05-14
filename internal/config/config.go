package config

import (
	"flag"
	"time"
)

type Config struct {
	CollectionInterval time.Duration
	OutputFormat       string
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

    	flag.Parse()

    	return Config{
    		CollectionInterval: *interval,
    		OutputFormat:       *output,
    	}
}