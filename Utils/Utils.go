package utils

import (
	"fmt"
	"time"

	dtos "github.com/adamwahyuh/fsd-golang-adamwh/Dtos"
	"github.com/adamwahyuh/fsd-golang-adamwh/config"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetSystemMetrics() (*dtos.ServerMetrics, error) {

	config.LoadEnv()
	key := config.GetEnv("GOLANG_KEY")

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, fmt.Errorf("gagal memuat timezone: %v", err)
	}

	now := time.Now().In(loc)
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
		Timestamp:    now.Format(time.RFC3339),
		CPUUsagePct:  cpuUsage,
		RAMTotalMB:   vMem.Total / 1024 / 1024,
		RAMUsedMB:    vMem.Used / 1024 / 1024,
		RAMUsagePct:  vMem.UsedPercent,
		DiskTotalGB:  diskStat.Total / 1024 / 1024 / 1024,
		DiskUsedGB:   diskStat.Used / 1024 / 1024 / 1024,
		DiskUsagePct: diskStat.UsedPercent,
		Key:          key,
	}

	return metrics, nil
}
