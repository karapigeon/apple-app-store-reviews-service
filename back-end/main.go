package main

import (
    "fmt"
    "log"
	"encoding/json"
	"strings"
	"io"
    "net/http"
	"os"
	// "time"
)

// Initialization function to start web server.
func main() {
    http.HandleFunc("/", rootHandler)
	http.HandleFunc("/data", dataHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Root handler to direct usage.
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Golang! Use the /data endpoint.")
}

// Constant for specific app ID in the App Store this service is observing.
const appId = "595068606"
// Constant for number of hours to filter array for.
const filterLastNumberHours = 72

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
	defer response.Body.Close()

	// Marshal data into inbound data type.
	var feedContainer InboundRSSFeedContainer
    err = json.NewDecoder(response.Body).Decode(&feedContainer)
    if err != nil {
        panic(err)
    }
    
	// Convert inbound data type to selective values for outbound data type.
	var outboundContainer []OutboundRSSFeedEntry
	for _, element := range feedContainer.Feed.Entry { 
		outboundContainer = append(outboundContainer, OutboundRSSFeedEntry{Content: element.Content.Label, Author: element.Author.Name.Label, Score: element.ImRating.Label, Timestamp: element.Updated.Label, Id: fmt.Sprintf("%s-%s", strings.ReplaceAll(element.Author.Name.Label, " ", ""), element.Updated.Label)})
    } 

	filename := fmt.Sprintf("%s.json", appId)

	// Open the file in read mode, create if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var localData []OutboundRSSFeedEntry
	err = json.NewDecoder(file).Decode(&localData)
	if err != nil {
		if err == io.EOF {
			// If file is empty, initialize the data
			localData = []OutboundRSSFeedEntry{}
		} else {
			fmt.Printf("Error decoding JSON: %v\n", err)
			return
		}
	}

	// Create a map to store the Entry structs indexed by the Id field
	entriesByUniqueId := make(map[string]OutboundRSSFeedEntry)

	// Populate the map
	for _, entry := range localData {
		entriesByUniqueId[entry.Id] = entry
	}

	for _, entry := range outboundContainer {
		// Check if the Id exists in the map
		if _, ok := entriesByUniqueId[entry.Id]; ok {
			continue
		} else {
			fmt.Printf("No matching record exists.")
			localData = append(localData, entry)
		}
	}

	jsonData, err := json.Marshal(localData)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	
	// Write the updated data back to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("File successfully updated.")
}

// Go type for data set stored in local disk file.s
type OutboundRSSFeedEntry struct {
	Content   string `json:"content"`
	Author    string `json:"author"`
	Score     string `json:"score"`
	Timestamp string `json:"timestamp"`
	Id string `json:"id"`
}

// Auto generated Go type from sampled JSON using
// https://transform.tools/json-to-go
type InboundRSSFeedContainer struct {
	Feed struct {
		Author struct {
			Name struct {
				Label string `json:"label"`
			} `json:"name"`
			URI struct {
				Label string `json:"label"`
			} `json:"uri"`
		} `json:"author"`
		Entry []struct {
			Author struct {
				URI struct {
					Label string `json:"label"`
				} `json:"uri"`
				Name struct {
					Label string `json:"label"`
				} `json:"name"`
				Label string `json:"label"`
			} `json:"author"`
			Updated struct {
				Label string `json:"label"`
			} `json:"updated"`
			ImRating struct {
				Label string `json:"label"`
			} `json:"im:rating"`
			ImVersion struct {
				Label string `json:"label"`
			} `json:"im:version"`
			ID struct {
				Label string `json:"label"`
			} `json:"id"`
			Title struct {
				Label string `json:"label"`
			} `json:"title"`
			Content struct {
				Label      string `json:"label"`
				Attributes struct {
					Type string `json:"type"`
				} `json:"attributes"`
			} `json:"content"`
			Link struct {
				Attributes struct {
					Rel  string `json:"rel"`
					Href string `json:"href"`
				} `json:"attributes"`
			} `json:"link"`
			ImVoteSum struct {
				Label string `json:"label"`
			} `json:"im:voteSum"`
			ImContentType struct {
				Attributes struct {
					Term  string `json:"term"`
					Label string `json:"label"`
				} `json:"attributes"`
			} `json:"im:contentType"`
			ImVoteCount struct {
				Label string `json:"label"`
			} `json:"im:voteCount"`
		} `json:"entry"`
		Updated struct {
			Label string `json:"label"`
		} `json:"updated"`
		Rights struct {
			Label string `json:"label"`
		} `json:"rights"`
		Title struct {
			Label string `json:"label"`
		} `json:"title"`
		Icon struct {
			Label string `json:"label"`
		} `json:"icon"`
		Link []struct {
			Attributes struct {
				Rel  string `json:"rel"`
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"attributes"`
		} `json:"link"`
		ID struct {
			Label string `json:"label"`
		} `json:"id"`
	} `json:"feed"`
}