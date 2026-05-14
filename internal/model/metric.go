package model

import "time"

type HostMetrics struct {
	Timestamp   time.Time `json:"timestamp"`
	Hostname    string    `json:"hostname"`
	CPUPercent  float64   `json:"cpu_percent"`
	MemoryTotal uint64    `json:"memory_total_bytes"`
	MemoryUsed  uint64    `json:"memory_used_bytes"`
	MemoryUsage float64   `json:"memory_usage_percent"`
}