package main

import (
	"encoding/json"
	"os"
)

func loadData() UserData {
	b, err := os.ReadFile(dataFile)
	if err != nil {
		return UserData{}
	}
	var d UserData
	json.Unmarshal(b, &d)
	return d
}

func saveData(d UserData) {
	b, _ := json.MarshalIndent(d, "", "  ")
	os.WriteFile(dataFile, b, 0644)
}
