# Drone Relay Agent

Lightweight host telemetry agent for the Drone Relay observability platform.

`drone-relay-agent` collects machine-level metrics such as CPU and memory usage,
batches telemetry locally, and streams data to the Drone Relay ingestion platform.

The project is part of a larger distributed observability and realtime event
processing system focused on:
- realtime telemetry ingestion
- distributed stream processing
- metrics aggregation
- large-scale observability
- live dashboards
- anomaly detection
- distributed tracing

---

# Current Features

- CPU usage collection
- Memory usage collection
- Host metadata collection
- Configurable collection interval
- JSON telemetry output
- Graceful shutdown handling

---

# Architecture

```text
+-------------------+
| drone-relay-agent|
|-------------------|
| CPU collector     |
| Memory collector  |
| Host metadata     |
+-------------------+
          |
          v
+-------------------+
| JSON Output       |
| stdout            |
+-------------------+
```

Future architecture:

```text
+-------------------+
|drone-relay-agent |
+-------------------+
          |
          v
+-------------------+
| Ingestion Gateway |
+-------------------+
          |
          v
+-------------------+
|       Kafka       |
+-------------------+
          |
          v
+-------------------+
| Stream Processors |
+-------------------+
          |
          v
+-------------------+
| Time-Series DB    |
+-------------------+
```

---

# Requirements

- Go 1.24+

---

# Installation

Clone repository:

```bash
git clone https://github.com/drone-relay/drone-relay-agent.git
cd drone-relay-agent
```

Install dependencies:

```bash
go mod tidy
```

---

# Running

Run with default settings:

```bash
go run ./cmd/drone-relay-agent
```

Run with custom interval:

```bash
go run ./cmd/drone-relay-agent --interval=2s
```

Run with explicit output format:

```bash
go run ./cmd/drone-relay-agent \
  --interval=2s \
  --output=json
```

---

# Example Output

```json
{
  "timestamp": "2026-05-14T12:00:00Z",
  "hostname": "node-01",
  "cpu_percent": 14.2,
  "memory_total_bytes": 17179869184,
  "memory_used_bytes": 8429314048,
  "memory_usage_percent": 49.06
}
```

---

# Repository Structure

```text
cmd/
  drone-relay-agent/
    main.go

internal/
  collector/
  config/
  model/
  output/
```

---

# Roadmap

## v0
- [x] Host metrics collection
- [x] JSON output
- [x] Configurable intervals

## v1
- [ ] Metric batching
- [ ] Local retry buffer
- [ ] Structured logging
- [ ] Config file support

## v2
- [ ] HTTP transport
- [ ] gRPC transport
- [ ] Compression
- [ ] Authentication

## v3
- [ ] Kafka/Redpanda ingestion
- [ ] Distributed ingestion gateway
- [ ] Backpressure handling

## v4
- [ ] Container metrics
- [ ] Kubernetes integration
- [ ] Distributed tracing
- [ ] OpenTelemetry compatibility

---

# Design Goals

drone-relay Agent is designed to be:
- lightweight
- resilient
- low-overhead
- horizontally scalable
- operationally simple

The project intentionally starts with a minimal feature set and incrementally
adds distributed systems capabilities over time.

---

# License

MIT