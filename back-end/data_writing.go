package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// DOC: Overwrite the local file with the revised/diffed dataset.
func OverwriteFileWithUpdatedDataSet(file *os.File, records []ReviewRecord) bool {
	// Encodes the data back into JSON.
	jsonData, err := json.Marshal(records)
	if err != nil {
		fmt.Printf("Error: Unable to marshal JSON. Error: %v", err)
		return false
	}

	// Attempts to overwrite the file using this new data collection.
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error: Unable to over/write file. Error: %v", err)
		return false
	}

	return true
}
