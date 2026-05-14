package config

import "time"

type Config struct {
	CollectionInterval time.Duration
}

func Default() Config {
	return Config{
		CollectionInterval: 5 * time.Second,
	}
}