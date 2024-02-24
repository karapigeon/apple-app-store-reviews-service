package main

import (
    "fmt"
	"strings"
)

func TransformForeignEntriesIntoRecords(container FeedContainer) []ReviewRecord {
	// Convert inbound data type to selective values for outbound data type.
	var records []ReviewRecord
	for _, element := range container.Feed.Entry { 
		records = append(records, ReviewRecord{Content: element.Content.Label, Author: element.Author.Name.Label, Score: element.ImRating.Label, Timestamp: element.Updated.Label, Id: fmt.Sprintf("%s-%s", strings.ReplaceAll(element.Author.Name.Label, " ", ""), element.Updated.Label)})
    } 

	return records
}

func FilterIncomingRecordsAgainstExistingLocalRecords(incoming []ReviewRecord, local []ReviewRecord) []ReviewRecord {
	// Create a map to store the Entry structs indexed by the Id field
	recordsById := make(map[string]ReviewRecord)

	// Populate the map
	for _, record := range local {
		recordsById[record.Id] = record
	}

	for _, record := range incoming {
		// Check if the Id exists in the map
		if _, ok := recordsById[record.Id]; ok {
			fmt.Printf("A record with id: %s already exists in the local file. Skipping", record.Id)
			continue
		} else {
			local = append(local, record)
		}
	}

	return local
}