# apple-app-store-reviews-service

Simple microservice to observe updates to Apple's App Store reviews RSS feed. Back-end will be written in Go, front-end will be written in JavaScript/React.

# Dev Setup

Clone the repo and run `$ cd back-end && go run *.go` from one terminal session and `$ cd front-end && npm run dev` in another. The back-end app will be available on http://localhost:5050, front-end app on http://localhost:3000. This is vaguely clunky but service orchestration was not part of the problem statement.

# Versions Used

* Go 1.21.7
* Node 20.11.0
