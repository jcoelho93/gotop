package stats

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReader(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	mockData := "Line 1\nLine 2\nLine 3\n"
	_, err = tempFile.WriteString(mockData)
	assert.NoError(t, err)
	err = tempFile.Close()
	assert.NoError(t, err)

	lines, err := ReadFileLines(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, []string{"Line 1", "Line 2", "Line 3"}, lines)
}
