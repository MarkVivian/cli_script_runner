package internal

import (
	"fmt"
	"os"
	"encoding/json"
)

func getJsonData() []JsonData {
	if debugMode {
		fmt.Println("Fetching JSON data...")
	}

	data, err := os.ReadFile("manifests.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil
	}

	var jsonData []JsonData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return jsonData
}

func buildIndex() map[string]map[string]JsonData {
	jsonData := getJsonData()
	index_map := make(map[string]map[string]JsonData)
	for _, item := range jsonData {
		if index_map[item.Platform] == nil {
			index_map[item.Platform] = make(map[string]JsonData)
		}
		index_map[item.Platform][item.ID] = item
	}


	return index_map
}
