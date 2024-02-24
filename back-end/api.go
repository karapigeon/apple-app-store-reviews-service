package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DOC: A constant for the (App Store) app ID this service is observing.
const appId = "595068606"

// DOC: Request handler for /data endpoint.
func DataHandler(w http.ResponseWriter, r *http.Request) {
	endpointUri := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=1/json", appId)
	response, err := http.Get(endpointUri)

	/*
		| This is one of these where this specific resource (Apple's RSS feed) is very predictable.
		| I tried to break it with weird appId values, etc and err was always nil because the RSS feed
		| handled any errors. I've made a general error capture so if it did it handle it gracefully.
	*/
	if err != nil || response.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Uncaught exception during request response"))
	}

	defer response.Body.Close()

	// DOC: Decode HTTP response body into inbound Go type.
	var feedContainer FeedContainer
	err = json.NewDecoder(response.Body).Decode(&feedContainer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400: Foreign JSON data could not be decoded."))
	}

	// DOC: Calls function in data_transform.go to transform FeedContainer into a collection
	// DOC: of ReviewRecord which is a Go type with only the fields we want to keep.
	incomingRecords := TransformForeignEntriesIntoRecords(feedContainer)

	// DOC: Calls function in data_reading.go to open the .json file on disk (or create it)
	// DOC: and if it has data, decode it into a collection of ReviewRecord.
	fileName := fmt.Sprintf("%s.json", appId)
	file, currentLocalRecords := ReadFileFromDiskWithFileNameAndReturnRecords(fileName)

	// DOC: Calls function in data_transform.go to filter incoming records against the existing
	// DOC: collection to reduce the I/O load/computational complexity using a map.
	proposedLocalRecords := FilterIncomingRecordsAgainstExistingLocalRecords(incomingRecords, currentLocalRecords)

	// DOC: Calls function in data_writing.go to overwrite the .json file with the new collection.
	result := OverwriteFileWithUpdatedDataSet(file, proposedLocalRecords)

	if result {
		lenDiff := len(proposedLocalRecords) - len(currentLocalRecords)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("200: %s new records written to file: %s on disk.", lenDiff, fileName)))
	} else {
		/*
			| This could have better error handling however I am handling the individual errors inside
			| the inner functions and they are mostly non-blocking due to the RSS feed architecture.
			| RSS feeds, especially Apple's resource handles errors very well internally.
		*/
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("418: I'm a teapot."))
	}

	defer file.Close()
}

// Root handler to direct usage.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang! Call the /data endpoint.")
}
