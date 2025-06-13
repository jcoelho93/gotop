package stats

// Tests for stats/memory.go

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStatsLoad(t *testing.T) {
	// Create a temporary file with mock memory stats
	tempFile, err := os.CreateTemp("", "meminfo")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write mock data to the file
	mockData := "MemTotal:       16384 kB\n" +
		"MemFree:        2048 kB\n" +
		"MemAvailable:   8192 kB\n"
	_, err = tempFile.WriteString(mockData)
	assert.NoError(t, err)

	// Close the file to flush the data
	err = tempFile.Close()
	assert.NoError(t, err)

	// Create a MemoryStats instance and load the data
	var memStats MemoryStats
	err = memStats.Load(tempFile.Name())
	assert.NoError(t, err)

	// Verify the loaded values
	assert.Equal(t, uint64(16384), memStats.Total)
	assert.Equal(t, uint64(2048), memStats.Free)
	assert.Equal(t, uint64(8192), memStats.Available)
}
