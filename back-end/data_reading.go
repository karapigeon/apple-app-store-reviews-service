package main

import (
    "fmt"
	"encoding/json"
	"io"
	"os"
)

// DOC: Public function that combines the two internal functions below.
func ReadFileFromDiskWithFileNameAndReturnRecords(fileName string) (*os.File, []ReviewRecord) {
	file := readFileFromLocalStorageWithFileName(fileName)
	records := decodeRecordsFromFile(file)

	return file, records
}

// DOC: This will attempt to open the file or create it in place if it does not exist.
func readFileFromLocalStorageWithFileName(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error: Unable to open or create file: %v. Investigate file system permissions.", err)
		return nil
	}

	return file
}

// DOC: This will attempt to decode the data from the file and unmarshal it into a collection of ReviewRecord.
func decodeRecordsFromFile(file *os.File) []ReviewRecord {
	var records []ReviewRecord
	err := json.NewDecoder(file).Decode(&records)
	if err != nil {
		// If file is empty, initialize the data
		if err == io.EOF {
			fmt.Printf("Error: The file was empty. Error: %v", err)
		} else {
			fmt.Printf("Error: The file could not be decoded. Error: %v", err)
		}

		// Create an empty array.
		records = []ReviewRecord{}
	}

	return records
}