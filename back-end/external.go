package main

import (
	"fmt"
	"log"
	"net/http"
)

// DOC: Port the app will be served on.
const appPort = "5050"

// DOC: Main initializer to fire up web server.
func main() {
	// Register a dummy root endpoint and app /data endpoint defined in api.go.
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/data", DataHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", appPort), nil))
}
