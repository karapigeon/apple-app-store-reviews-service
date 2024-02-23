package main

import (
    "fmt"
    "log"
	"os"
	"io/ioutil"
    "net/http"
)

// Initialization function to start web server.
func main() {
    http.HandleFunc("/", rootHandler)
	http.HandleFunc("/data", dataHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Root handler to direct usage.
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Golang! Use the /rss endpoint.")
}

// App Store specific app ID this service is observing.
const appId = "595068606"

// HTTP request handler for /data endpoint.
func dataHandler(w http.ResponseWriter, r *http.Request) {
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

	// Unwrap the response body.
	responseBody, err := ioutil.ReadAll(response.Body)

	// Handle generalized error type for response body object.
	// TODO: This could be expanded to handle specific scenarios.
	if err != nil {
		fmt.Printf("Error: Client was unable to read the response body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("500 - Uncaught exception while reading the response body"))
	}

	fmt.Printf("Info: Client has response body: %s", responseBody)
}