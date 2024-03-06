package kafka

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Read(path string) (*kafka.ConfigMap, error) {
	config := &kafka.ConfigMap{}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			before, after, found := strings.Cut(line, "=")
			if found {
				parameter := strings.TrimSpace(before)
				value := strings.TrimSpace(after)
				config.SetKey(parameter, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to read: %s", err)
	}

	return config, nil
}
