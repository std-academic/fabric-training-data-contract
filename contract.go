package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// I think this is "Create or Update" function
func (s *SmartContract) CreateData(ctx contractapi.TransactionContextInterface, id string, data string, metadataJson string) error {
	var metadata Metadata
	if err := json.Unmarshal([]byte(metadataJson), &metadata); err != nil {
		return fmt.Errorf("invalid metadata format: %s", err.Error())
	}

	trainingData := TrainingData{
		ID:       id,
		Data:     data,
		Metadata: metadata,
	}

	if err := validateTrainingData(trainingData); err != nil {
		return err
	}

	trainingDataAsBytes, _ := json.Marshal(trainingData)
	return ctx.GetStub().PutState(id, trainingDataAsBytes)
}

func (s *SmartContract) QueryData(ctx contractapi.TransactionContextInterface, id string) (*TrainingData, error) {
	dataAsBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if dataAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	trainingData := new(TrainingData)
	_ = json.Unmarshal(dataAsBytes, trainingData)

	return trainingData, nil
}

func (s *SmartContract) QueryMetadata(ctx contractapi.TransactionContextInterface, id string) (*Metadata, error) {
	dataAsBytes, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if dataAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	trainingData := new(TrainingData)
	_ = json.Unmarshal(dataAsBytes, trainingData)

	return &trainingData.Metadata, nil
}

func (s *SmartContract) QueryAllMetadatas(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		trainingData := new(TrainingData)
		_ = json.Unmarshal(queryResponse.Value, trainingData)

		queryResult := QueryResult{ID: trainingData.ID, Metadata: &trainingData.Metadata}
		results = append(results, queryResult)
	}

	return results, nil
}

func (s *SmartContract) ChangeMetadata(ctx contractapi.TransactionContextInterface, id string, metadataJson string) error {
	trainingData, err := s.QueryData(ctx, id)

	if err != nil {
		return err
	}

	var metadata Metadata
	if err := json.Unmarshal([]byte(metadataJson), &metadata); err != nil {
		return fmt.Errorf("invalid metadata format: %s", err.Error())
	}

	if err := validateMetadata(metadata); err != nil {
		return err
	}

	trainingData.Metadata = metadata

	dataAsBytes, _ := json.Marshal(trainingData)
	return ctx.GetStub().PutState(id, dataAsBytes)
}

func (s *SmartContract) ChangeData(ctx contractapi.TransactionContextInterface, id string, data string) error {
	trainingData, err := s.QueryData(ctx, id)

	if err != nil {
		return err
	}

	trainingData.Data = data

	dataAsBytes, _ := json.Marshal(trainingData)
	return ctx.GetStub().PutState(id, dataAsBytes)
}
