package stats

import (
	"fmt"
	"strconv"
	"strings"
)

type UptimeStats struct {
	Seconds float64
	Idle	float64
}

func (u *UptimeStats) Load(path string) error {
	lines, err := ReadFileLines(path)
	if err != nil {
		return fmt.Errorf("failed to read uptime file: %w", err)
	}
	if len(lines) != 1 {
		return fmt.Errorf("unexpected number of lines in uptime file: %d", len(lines))
	}

	parts := strings.Fields(lines[0])
	if len(parts) != 2 {
		return fmt.Errorf("unexpected format in uptime file: %s", lines[0])
	}

	uptime, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return fmt.Errorf("failed to parse uptime: %w", err)
	}
	idle, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return fmt.Errorf("failed to parse idle time: %w", err)
	}
	u.Seconds = uptime
	u.Idle = idle
	return nil
}
