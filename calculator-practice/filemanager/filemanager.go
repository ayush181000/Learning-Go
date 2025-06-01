package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, nil
}

func (fm FileManager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	time.Sleep(3 * time.Second)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert to json")
	}

	return nil
}
