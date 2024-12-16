package service

import (
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AIService struct {
	Client HTTPClient
}

func (s *AIService) AnalyzeData(table map[string][]string, query, token string) (string, error) {
	if len(table) == 0 {
		return "", errors.New("table data is empty")
	}
	if query == "" {
		return "", errors.New("query is required")
	}

	// Prepare the request payload
	payload := map[string]interface{}{
		"inputs": map[string]interface{}{
			"query": query,
			"table": table,
		},
	}

	// Serialize payload into JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to serialize payload: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := s.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	// Extract result (assuming the format is {"cells": ["result"]})
	if cells, ok := response["cells"].([]interface{}); ok && len(cells) > 0 {
		if result, ok := cells[0].(string); ok {
			return result, nil
		}
	}

	return "", errors.New("unexpected response format")
}

// func (s *AIService) ChatWithAI(context, query, token string) (model.ChatResponse, error) {
// 	if query == "" {
// 		return model.ChatResponse{}, errors.New("query is required")
// 	}

// 	// Prepare the payload
// 	payload := map[string]interface{}{
// 		"inputs": map[string]interface{}{
// 			"context": context,
// 			"query":   query,
// 		},
// 	}

// 	// Serialize payload to JSON
// 	payloadJSON, err := json.Marshal(payload)
// 	if err != nil {
// 		return model.ChatResponse{}, fmt.Errorf("failed to serialize payload: %v", err)
// 	}

// 	// Create HTTP request
// 	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/microsoft/DialoGPT-medium", bytes.NewBuffer(payloadJSON))
// 	if err != nil {
// 		return model.ChatResponse{}, fmt.Errorf("failed to create request: %v", err)
// 	}

// 	// Add headers
// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Content-Type", "application/json")

// 	// Execute the request
// 	resp, err := s.Client.Do(req)
// 	if err != nil {
// 		return model.ChatResponse{}, fmt.Errorf("failed to execute request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Check for HTTP errors
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return model.ChatResponse{}, fmt.Errorf("received non-OK status code: %d, body: %s", resp.StatusCode, string(body))
// 	}

// 	// Parse the response
// 	var response []model.ChatResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return model.ChatResponse{}, fmt.Errorf("failed to parse response: %v", err)
// 	}

// 	if len(response) == 0 {
// 		return model.ChatResponse{}, errors.New("empty response from AI model")
// 	}

// 	// Return the first response
// 	return response[0], nil
// }

func (s *AIService) ChatWithAI(context, query, token string) (model.ChatResponse, error) {
	if query == "" {
		return model.ChatResponse{}, errors.New("query is required")
	}

	// Prepare the payload
	payload := map[string]interface{}{
		"inputs": query,
		"parameters": map[string]interface{}{
			"max_tokens": 500,
			"stream":     false,
		},
	}

	// Serialize payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to serialize payload: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/microsoft/Phi-3.5-mini-instruct", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to create request: %v", err)
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := s.Client.Do(req)
	if err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return model.ChatResponse{}, fmt.Errorf("received non-OK status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var response []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.ChatResponse{}, fmt.Errorf("failed to parse response: %v", err)
	}

	// Check if there are any items in the array
	if len(response) > 0 {
		// Extract the content from the first item
		if content, ok := response[0]["generated_text"].(string); ok {
			return model.ChatResponse{
				GeneratedText: content,
			}, nil
		}
	}

	return model.ChatResponse{}, errors.New("no content found in response")
}
