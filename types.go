package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	Owner       string   `json:"owner"`
	Name        string   `json:"name"`
	Tasks       []string `json:"tasks"`
	Modalities  []string `json:"modalities"`
	Formats     []string `json:"formats"`
	Languages   []string `json:"languages"`
	Size        int64    `json:"size"`
	Tags        []string `json:"tags"`
	Libraries   []string `json:"libraries"`
	License     string   `json:"license"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	// Rows count
	Rows int64 `json:"rows"`
	// use RFC3339 format
	UploadDate string `json:"uploadDate"`
	// Content Identifier, the hash of the data, in case too big
	CID string `json:"cid"`
}

type TrainingData struct {
	// ID is the unique identifier of the training data
	// UUID format
	ID       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Data     string   `json:"data"`
}

type QueryResult struct {
	ID       string    `json:"id"`
	Metadata *Metadata `json:"metadata"`
}

func validateMetadata(metadata Metadata) error {
	if strings.TrimSpace(metadata.Owner) == "" {
		return fmt.Errorf("owner is required")
	}
	if strings.TrimSpace(metadata.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if metadata.Size <= 0 {
		return fmt.Errorf("size must be greater than zero")
	}
	if metadata.Version == "" {
		return fmt.Errorf("version is required")
	}
	if metadata.Rows <= 0 {
		return fmt.Errorf("rows must be greater than zero")
	}
	if _, err := time.Parse(time.RFC3339, metadata.UploadDate); err != nil {
		return fmt.Errorf("invalid uploadDate format")
	}

	return nil
}

func validateTrainingData(trainingData TrainingData) error {
	// check if ID is uuid
	if _, err := uuid.Parse(trainingData.ID); err != nil {
		return fmt.Errorf("invalid uuID format: %s", err.Error())
	}
	if err := validateMetadata(trainingData.Metadata); err != nil {
		return err
	}
	if trainingData.Data == "" && trainingData.Metadata.CID == "" {
		return fmt.Errorf("data or CID is required")
	}

	return nil
}
