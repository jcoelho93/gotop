package stats

import (
	"strconv"
	"strings"
)

type MemoryStats struct {
	Total uint64
	Free  uint64
	Available uint64
}

func (m *MemoryStats) Load(path string) error {
	lines, err := ReadFileLines(path)
	if err != nil {
		return err
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		key := fields[0]
		value, _ := strconv.ParseUint(fields[1], 10, 64)
		switch key {
		case "MemTotal:":
			m.Total = value
		case "MemFree:":
			m.Free = value
		case "MemAvailable:":
			m.Available = value
		default:
			continue
		}
	}
	return nil
}
