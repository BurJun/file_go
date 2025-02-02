package service

import (
	"bufio"
	"os"
)

type Producer interface {
	Produce() ([]string, error)
}

type FileProducer struct {
	filePath string
}

func NewFileProducer(filePath string) *FileProducer {
	return &FileProducer{filePath: filePath}
}

func (fp *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(fp.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
