package utils

import (
	"fmt"
	"time"

	dtos "github.com/adamwahyuh/fsd-golang-adamwh/Dtos"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetSystemMetrics() (*dtos.ServerMetrics, error) {
	// ambil data cpu
	cpuPercentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data CPU: %v", err)
	}
	cpuUsage := 0.0
	if len(cpuPercentages) > 0 {
		cpuUsage = cpuPercentages[0]
	}

	// ambil data RAM
	vMem, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data RAM: %v", err)
	}

	// ambil data Disk (Root directory "/")
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data Disk: %v", err)
	}

	metrics := &dtos.ServerMetrics{
		Timestamp:    time.Now().Format(time.RFC3339),
		CPUUsagePct:  cpuUsage,
		RAMTotalMB:   vMem.Total / 1024 / 1024,
		RAMUsedMB:    vMem.Used / 1024 / 1024,
		RAMUsagePct:  vMem.UsedPercent,
		DiskTotalGB:  diskStat.Total / 1024 / 1024 / 1024,
		DiskUsedGB:   diskStat.Used / 1024 / 1024 / 1024,
		DiskUsagePct: diskStat.UsedPercent,
	}

	return metrics, nil
}
