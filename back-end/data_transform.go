package main

import (
    "fmt"
	"strings"
)

// DOC: Transform the contents of the inbound data type to record data type.
// DOC: This only contains the fields required by the front-end.
func TransformForeignEntriesIntoRecords(container FeedContainer) []ReviewRecord {
	var records []ReviewRecord
	for _, element := range container.Feed.Entry { 
		// Picks out the content, author, score, and timestamp.
		// As well as creates a sanitized composit key of author+timestamp to ensure review uniqueness.
		records = append(records, ReviewRecord{Content: element.Content.Label, Author: element.Author.Name.Label, Score: element.ImRating.Label, Timestamp: element.Updated.Label, Id: fmt.Sprintf("%s-%s", strings.ReplaceAll(element.Author.Name.Label, " ", ""), element.Updated.Label)})
    } 

	return records
}

// DOC: Filters out records from the new collection that already exists in the current local collection.
func FilterIncomingRecordsAgainstExistingLocalRecords(incoming []ReviewRecord, local []ReviewRecord) []ReviewRecord {
	// Using a map because it's less computationally complex than a (nested) for-loop.
	recordsById := make(map[string]ReviewRecord)
	// Populates the map using current local collection.
	for _, record := range local {
		recordsById[record.Id] = record
	}

	for _, record := range incoming {
		// Checks if the id (new collection) already exists in the local collection.
		if _, ok := recordsById[record.Id]; ok {
			fmt.Printf("Info: A record with id: %s already exists in the local file. Skipping", record.Id)
			continue
		} else {
			local = append(local, record)
		}
	}

	return local
}