package main

// Go type for data set stored in local disk file.
type ReviewRecord struct {
	Content   string `json:"content"`
	Author    string `json:"author"`
	Score     string `json:"score"`
	Timestamp string `json:"timestamp"`
	Id string `json:"id"`
}