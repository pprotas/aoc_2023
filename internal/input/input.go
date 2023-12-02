package input

import (
	"fmt"
	"os"
	"path/filepath"

	"aoc_2023/internal/client"
)

const inputDir = "input"

func GetInput(day int) (string, error) {
	filePath := filepath.Join(inputDir, fmt.Sprintf("%02d.txt", day))

	input, err := InputFromFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return CacheInput(day, filePath)
		}
		return "", err
	}

	return input, nil
}

func InputFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CacheInput(day int, filePath string) (string, error) {
	err := os.MkdirAll(inputDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	input := client.FetchInput(day)

	err = os.WriteFile(filePath, []byte(input), 0o644)
	if err != nil {
		return "", err
	}

	return input, nil
}
