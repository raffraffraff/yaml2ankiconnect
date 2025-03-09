package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// YAML structure
type YAMLData struct {
	Models map[string]Model `yaml:"models"`
	Decks  map[string]Deck  `yaml:"decks"`
}

type Model struct {
	Fields    []string    `yaml:"fields"`
	Templates []Template `yaml:"templates"`
}

type Template struct {
	Name  string `yaml:"name"`
	Front string `yaml:"front"`
	Back  string `yaml:"back"`
}

type Deck struct {
	ModelName string `yaml:"modelName"`
	Cards     []Card `yaml:"cards"`
}

type Card map[string]interface{}

func processYAML(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var yamlData YAMLData
	err = yaml.Unmarshal(data, &yamlData)
	if err != nil {
		return err
	}

	// Create models in Anki
	for modelName, model := range yamlData.Models {
		err := createAnkiModel(modelName, model)
		if err != nil {
			fmt.Println("Error creating model:", err)
		}
	}

	// Process decks and cards
	for deckName, deck := range yamlData.Decks {
		for _, card := range deck.Cards {
			// Extract tags
			var tags []string
			if t, exists := card["tags"]; exists {
				if tagList, ok := t.([]interface{}); ok {
					for _, tag := range tagList {
						if tagStr, ok := tag.(string); ok {
							tags = append(tags, tagStr)
						}
					}
				}
			}
			delete(card, "tags")

			// Process media
			var mediaFiles []string
			if m, exists := card["media"]; exists {
				if mediaList, ok := m.([]interface{}); ok {
					for _, media := range mediaList {
						if mediaStr, ok := media.(string); ok {
							mediaFiles = append(mediaFiles, mediaStr)
							err := uploadMedia(mediaStr)
							if err != nil {
								fmt.Println("Error uploading media:", err)
							}
						}
					}
				}
			}
			delete(card, "media")

			// Replace media filenames in fields
			for key, value := range card {
				if strVal, ok := value.(string); ok {
					for _, media := range mediaFiles {
						if filepath.Ext(media) == ".mp3" {
							strVal = fmt.Sprintf("[sound:%s]", filepath.Base(media))
						} else {
							strVal = fmt.Sprintf("<img src=\"%s\">", filepath.Base(media))
						}
					}
					card[key] = strVal
				}
			}

			// Create Anki note
			note := Note{
				DeckName:  deckName,
				ModelName: deck.ModelName,
				Fields:    card,
				Tags:      tags,
			}

			request := AnkiRequest{
				Action:  "addNote",
				Version: 6,
				Params:  AddNoteParams{Note: note},
			}

			response, err := sendAnkiRequest(request)
			if err != nil {
				fmt.Printf("Error creating %s card in deck %s: %s\n", deck.ModelName, deckName, err)
			} else {
				fmt.Printf("Card added to deck %s: %s\n", deckName, response)
			}
		}
	}

	return nil
}
