package main

import (
	"fmt"
	"strings"
	"sort"
	"time"
)

// DOC: Transform the contents of the inbound data type to record data type.
// DOC: This only contains the fields required by the front-end.
func TransformForeignEntriesIntoRecords(container FeedContainer) []ReviewRecord {
	var records []ReviewRecord
	for _, element := range container.Feed.Entry {
		unixTimestamp := transformRFC3339TimestampIntoMilliseconds(element.Updated.Label)

		if unixTimestamp != nil {
			// Picks out the content, author, score, and timestamp.
			// As well as creates a sanitized composit key of author+timestamp to ensure review uniqueness.
			records = append(records, ReviewRecord{Content: element.Content.Label, Author: element.Author.Name.Label, Score: element.ImRating.Label, Timestamp: *unixTimestamp, Id: fmt.Sprintf("%s-%s", strings.ReplaceAll(element.Author.Name.Label, " ", ""), *unixTimestamp)})
		}
	}

	return records
}

func transformRFC3339TimestampIntoMilliseconds(timestampString string) *string {
	 // Parse the timestamp string into a time.Time object.
	 timestamp, err := time.Parse(time.RFC3339, timestampString)
	 if err != nil {
		 fmt.Println("Error: Unable to parse timestamp as RFC3339. Error: %v", err)
		 return nil
	 }

	 unixTimestamp := timestamp.UnixNano() / int64(time.Millisecond)
	 castedTimestamp := fmt.Sprint(unixTimestamp)

	 // Convert time to milliseconds.
	 return &castedTimestamp
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

	//  Sort the array of records by timestamp in descending order.
	// Sort the array of records by timestamp in descending order
    sort.Slice(local, func(i, j int) bool {
        timeI, err := time.Parse(time.RFC3339, local[i].Timestamp)
        if err != nil {
            return false
        }
        timeJ, err := time.Parse(time.RFC3339, local[j].Timestamp)
        if err != nil {
            return false
        }
        return timeI.After(timeJ)
    })

	return local
}
