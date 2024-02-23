package main

import (
    "fmt"
    "log"
	"encoding/json"
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
    fmt.Fprintf(w, "Hello from Golang! Use the /data endpoint.")
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
	defer response.Body.Close()

	// Marshal data into inbound data type.
	var feedContainer InboundRSSFeedContainer
    err = json.NewDecoder(response.Body).Decode(&feedContainer)
    if err != nil {
        panic(err)
    }
    
	// Convert inbound data type to selective values for outbound data type.
	var outboundContainer []OutboundContainerEntry
	for _, element := range feedContainer.Feed.Entry { 
		outboundContainer = append(outboundContainer, OutboundContainerEntry{Content: element.Content.Label, Author: element.Author.Name.Label, Score: element.ImRating.Label, Timestamp: element.Updated.Label})
    } 

	fmt.Printf("Info: Outbound container: %s", outboundContainer)
}

// Go type for data set stored in local disk file.s
type OutboundContainerEntry struct {
	Content   string `json:"content"`
	Author    string `json:"author"`
	Score     string `json:"score"`
	Timestamp string `json:"timestamp"`
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