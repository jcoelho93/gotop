package stats

import (
	"os"
	"strings"
)

type ProcFileParser interface {
	Load(path string) error
}

func ReadFileLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines, nil
}
