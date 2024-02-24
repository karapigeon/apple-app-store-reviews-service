package main

import (
    "fmt"
	"encoding/json"
	"io"
	"os"
)

func ReadFileFromDiskWithFileNameAndReturnRecords(fileName string) (*os.File, []ReviewRecord) {
	file := readFileFromLocalStorageWithFileName(fileName)
	records := decodeRecordsFromFile(file)

	return file, records
}

func readFileFromLocalStorageWithFileName(fileName string) *os.File {
	// Open the file in read mode, create if it doesn't exist
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}
	// defer file.Close()

	return file
}

func decodeRecordsFromFile(file *os.File) []ReviewRecord {
	var records []ReviewRecord
	err := json.NewDecoder(file).Decode(&records)
	if err != nil {
		if err == io.EOF {
			// If file is empty, initialize the data
			records = []ReviewRecord{}
		} else {
			fmt.Printf("Error decoding JSON: %v\n", err)
		}
	}

	return records
}