package dtos

type ServerMetrics struct {
	Timestamp    string  `json:"timestamp"`
	CPUUsagePct  float64 `json:"cpu_usage_pct"`
	RAMTotalMB   uint64  `json:"ram_total_mb"`
	RAMUsedMB    uint64  `json:"ram_used_mb"`
	RAMUsagePct  float64 `json:"ram_usage_pct"`
	DiskTotalGB  uint64  `json:"disk_total_gb"`
	DiskUsedGB   uint64  `json:"disk_used_gb"`
	DiskUsagePct float64 `json:"disk_usage_pct"`
	Key          string  `json:"key"`
}
