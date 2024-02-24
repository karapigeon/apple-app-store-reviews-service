package main

import (
    "fmt"
	"encoding/json"
    "net/http"
)

// Constant for specific app ID in the App Store this service is observing.
const appId = "595068606"

// HTTP request handler for /data endpoint.
func DataHandler(w http.ResponseWriter, r *http.Request) {
	// Add appId to request URL.
    requestUrl := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=1/json", appId)
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	
	// Configure ResponseWriter to gracefully exit. 
	// This would only really happen if the string interpolation for appId failed.
	if err != nil {
		fmt.Printf("Error: Client could not create request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte(fmt.Sprintf("500 - Could not create request with %s", requestUrl)))
	}

	// Launch request and get back response.
	response, err := http.DefaultClient.Do(request)

	// Handle generalized error type for response object.
	// TODO: This could be expanded to handle specific status codes/scenarios.
	if err != nil {
		fmt.Printf("Error: Client produced error when making http request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("500 - Uncaught exception during request response"))
	}

	fmt.Printf("Info: Client returned status code: %s with response", response.StatusCode)
	defer response.Body.Close()

	// Marshal data into inbound data type.
	var feedContainer FeedContainer
    err = json.NewDecoder(response.Body).Decode(&feedContainer)
    if err != nil {
        panic(err)
    }
    
	incomingRecords := TransformForeignEntriesIntoRecords(feedContainer)

	fileName := fmt.Sprintf("%s.json", appId)
	file, currentLocalRecords := ReadFileFromDiskWithFileNameAndReturnRecords(fileName)
	proposedLocalRecords := FilterIncomingRecordsAgainstExistingLocalRecords(incomingRecords, currentLocalRecords)
	result := RewriteFileWithUpdatedDataSet(file, proposedLocalRecords)

	if result {
		lenDiff := len(proposedLocalRecords) - len(currentLocalRecords)
		fmt.Printf("Info: %s was successfully overwritten with %s new records.", fileName, lenDiff)
	} else {
		fmt.Printf("Info: Something went wrong.")
	}

	
}

// Root handler to direct usage.
func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Golang! Call the /data endpoint.")
}