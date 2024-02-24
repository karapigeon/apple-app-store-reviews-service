package main

import (
    "log"
    "net/http"
)

// Initialization function to start web server.
func main() {
    http.HandleFunc("/", RootHandler)
	http.HandleFunc("/data", DataHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO
// Constant for number of hours to filter array for.
// const filterLastNumberHours = 72