package service

import (
	repository "a21hc3NpZ25tZW50/repository/fileRepository"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type FileService struct {
	Repo *repository.FileRepository
}

func (s *FileService) ProcessFile(filename string) (map[string][]string, error) {
	var fileContent []byte
	var err error

	// Check if filename is a file or direct content
	if _, statErr := os.Stat(filename); statErr == nil {
		fileContent, err = s.Repo.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}
	} else {
		// Treat filename as raw content if file doesn't exist
		fileContent = []byte(filename)
	}

	if len(fileContent) == 0 {
		return nil, errors.New("file content is empty")
	}

	// Parse CSV content
	reader := csv.NewReader(bytes.NewReader(fileContent))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV content: %v", err)
	}

	if len(records) == 0 {
		return nil, errors.New("CSV file has no records")
	}

	// Extract headers from the first row
	headers := records[0]
	if len(headers) == 0 {
		return nil, errors.New("CSV header row is empty")
	}

	// Initialize result map
	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	// Populate result map with data rows
	for _, row := range records[1:] {
		for i, value := range row {
			if i < len(headers) {
				result[headers[i]] = append(result[headers[i]], value)
			}
		}
	}

	return result, nil
}
