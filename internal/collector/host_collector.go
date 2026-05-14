package collector

import (
	"context"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"

	"github.com/drone-relay/drone-relay-agent/internal/model"
)

type HostCollector struct {
	hostname string
}

func NewHostCollector() (*HostCollector, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &HostCollector{
		hostname: hostname,
	}, nil
}

func (c *HostCollector) Collect(ctx context.Context) (model.HostMetrics, error) {
	cpuPercentages, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil {
		return model.HostMetrics{}, err
	}

	memoryStats, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return model.HostMetrics{}, err
	}

	var cpuPercent float64
	if len(cpuPercentages) > 0 {
		cpuPercent = cpuPercentages[0]
	}

	return model.HostMetrics{
		Timestamp:   time.Now().UTC(),
		Hostname:    c.hostname,
		CPUPercent:  cpuPercent,
		MemoryTotal: memoryStats.Total,
		MemoryUsed:  memoryStats.Used,
		MemoryUsage: memoryStats.UsedPercent,
	}, nil
}