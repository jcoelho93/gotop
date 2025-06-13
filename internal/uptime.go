package uptime

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Uptime struct {
	Seconds float64
	Idle	float64
}

type UptimeReader struct {
	path string
}

func NewUptimeReader(path string) *UptimeReader {
	return &UptimeReader{path: path}
}

func (r *UptimeReader) Read() (*Uptime, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read uptime file: %w", err)
	}

	parts := strings.Fields(string(data))
	if len(parts) < 2 {
		return nil, fmt.Errorf("unexpected format in uptime file: %s", string(data))
	}

	uptime, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uptime: %w", err)
	}
	idle, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse idle time: %w", err)
	}

	return &Uptime{
		Seconds: uptime,
		Idle: idle,
	}, nil
}
