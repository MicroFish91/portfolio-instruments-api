package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := ".thunderclient/Portfolio_Instruments_API_Dev.json"

	// Read
	readBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var tc_requests ThunderClientRequests
	json.Unmarshal(readBytes, &tc_requests)

	// Scrub tokens
	for i, req := range tc_requests.Requests {
		if req.Auth != nil && req.Auth.Bearer != "" {
			tc_requests.Requests[i].Auth.Bearer = ""
		} else {
			tc_requests.Requests[i].Auth = nil
		}
	}

	// Write
	writeBytes, err := json.Marshal(&tc_requests)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, writeBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully executed scrubbing of thunderclient json data!")
}

// ThunderClient Types
type ThunderClientRequests struct {
	ClientName     string                 `json:"clientName"`
	CollectionName string                 `json:"collectionName"`
	CollectionId   string                 `json:"collectionId"`
	DateExported   string                 `json:"dateExported"`
	Version        string                 `json:"version"`
	Folders        []ThunderClientFolder  `json:"folders"`
	Requests       []ThunderClientRequest `json:"requests"`
	Ref            string                 `json:"ref"`
}

type ThunderClientFolder struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	ContainerId string `json:"containerId"`
	Created     string `json:"created"`
	SortNum     int    `json:"sortNum"`
}

type ThunderClientRequest struct {
	Id          string             `json:"_id"`
	ColId       string             `json:"colId"`
	ContainerId string             `json:"containerId"`
	Name        string             `json:"name"`
	Url         string             `json:"url"`
	Method      string             `json:"method"`
	SortNum     int                `json:"sortNum"`
	Created     string             `json:"created"`
	Modified    string             `json:"modified"`
	Headers     []string           `json:"headers"`
	Auth        *ThunderClientAuth `json:"auth"`
	Body        any                `json:"body"`
	Params      any                `json:"params"`
}

type ThunderClientAuth struct {
	Type   string `json:"type"`
	Bearer string `json:"bearer"`
}
