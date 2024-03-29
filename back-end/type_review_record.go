package main

// DOC: Go type for interfacing with records stored in local disk file.
type ReviewRecord struct {
	Content   string `json:"content"`
	Author    string `json:"author"`
	Score     string `json:"score"`
	Timestamp int64 `json:"timestamp"`
	Id        string `json:"id"`
}
