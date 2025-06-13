package display

import (
	"fmt"
	"gotop/internal/pkg/stats"
)

type Top struct {}

func (t *Top) Display() error {

	uptimeStats := stats.UptimeStats{}
	err := uptimeStats.Load("/proc/uptime")
	if err != nil {
		fmt.Println("Error loading uptime stats:", err)
		return err
	}
	memoryStats := stats.MemoryStats{}
	err = memoryStats.Load("/proc/meminfo")
	if err != nil {
		fmt.Println("Error loading memory stats:", err)
		return err
	}

	fmt.Printf("\rUptime: %.2f seconds, Idle: %.2f seconds\n", uptimeStats.Seconds, uptimeStats.Idle)
	fmt.Printf("\rMemory Total: %d, Free: %d", memoryStats.Total, memoryStats.Free)
	fmt.Print("\033[1A")
	return nil
}
