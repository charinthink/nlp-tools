package utilities

import (
	"bufio"
	"os"
)

func LoadDict(corpusDir string) ([]string, error) {
	file, err := os.Open(corpusDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
