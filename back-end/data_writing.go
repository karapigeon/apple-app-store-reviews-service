package main

import (
    "fmt"
	"encoding/json"
	"os"
)

func RewriteFileWithUpdatedDataSet(file *os.File, records []ReviewRecord) bool {
	jsonData, err := json.Marshal(records)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return false
	}
	
	// Write the updated data back to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return false
	}

	fmt.Println("File successfully updated.")

	return true
}