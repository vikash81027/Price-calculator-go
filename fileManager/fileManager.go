package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() (*[] string, error) {
	file, error := os.Open(fm.InputFilePath)

	if error != nil {
		file.Close()
		return nil, errors.New("failed to open the file")
	}
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	error = scanner.Err()

	if error != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return &lines, nil
}

func (fm *FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath);

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to connect data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{InputFilePath: inputPath, OutputFilePath: outputPath}
}