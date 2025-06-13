package stats

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUptimeStatsLoad(t *testing.T) {
	tempFile, err := os.CreateTemp("", "uptime_test")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	mockData := "12345.67 8901.23\n"
	_, err = tempFile.WriteString(mockData)
	assert.NoError(t, err)
	err = tempFile.Close()
	assert.NoError(t, err)

	var uptimeStats UptimeStats
	err = uptimeStats.Load(tempFile.Name())
	assert.NoError(t, err)

	assert.Equal(t, 12345.67, uptimeStats.Seconds)
	assert.Equal(t, 8901.23, uptimeStats.Idle)
}
