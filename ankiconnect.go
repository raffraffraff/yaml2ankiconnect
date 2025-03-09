package main

import (
	"errors"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// AnkiConnect structures
type AnkiRequest struct {
	Action  string      `json:"action"`
	Version int         `json:"version"`
	Params  interface{} `json:"params,omitempty"`
}

type Note struct {
	DeckName  string                 `json:"deckName"`
	ModelName string                 `json:"modelName"`
	Fields    map[string]interface{} `json:"fields"`
	Tags      []string               `json:"tags,omitempty"`
}

type AddNoteParams struct {
	Note Note `json:"note"`
}

// Function to send JSON requests to AnkiConnect
func sendAnkiRequest(payload AnkiRequest) (map[string]interface{}, error) {
	url := "http://localhost:8765"
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// this just catches errors when trying to post
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        	return nil, fmt.Errorf("failed to decode response: %w", err)
    	}

	res := result["result"]
	if res == nil {
		errorText := result["error"].(string)
		return nil, errors.New(errorText)
	}
	return result, nil

}

// Function to check if model exists
func modelExists(modelName string) bool {
	request := AnkiRequest{
		Action:  "modelNames",
		Version: 6,
	}
	response, err := sendAnkiRequest(request)
	if err != nil {
		fmt.Println("Error checking model:", err)
		return false
	}

	if models, ok := response["result"].([]interface{}); ok {
		for _, model := range models {
			if model.(string) == modelName {
				return true
			}
		}
	}
	return false
}

// Function to create a new Anki model
func createAnkiModel(name string, model Model) error {
	if modelExists(name) {
		fmt.Println("Model already exists:", name)
		return nil
	}

	ankiModel := map[string]interface{}{
		"modelName": name,
		"inOrderFields": model.Fields,
		"css":          "", // You can add CSS here
		"cardTemplates": []map[string]string{},
	}

	for _, tmpl := range model.Templates {
		ankiModel["cardTemplates"] = append(ankiModel["cardTemplates"].([]map[string]string), map[string]string{
			"Name":  tmpl.Name,
			"Front": tmpl.Front,
			"Back":  tmpl.Back,
		})
	}

	request := AnkiRequest{
		Action:  "createModel",
		Version: 6,
		Params:  ankiModel,
	}

	_, err := sendAnkiRequest(request)
	return err
}

// Function to upload media files
func uploadMedia(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	filename := filepath.Base(filePath)

	request := AnkiRequest{
		Action:  "storeMediaFile",
		Version: 6,
		Params: map[string]string{
			"filename": filename,
			"data":     string(data),
		},
	}

	_, err = sendAnkiRequest(request)
	return err
}
