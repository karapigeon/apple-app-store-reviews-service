# apple-app-store-reviews-service

Simple microservice to observe updates to Apple's App Store reviews RSS feed. Back-end will be written in Go, front-end will be written in TypeScript/React/Remix.

# Summary

# Notes

* I was originally going to write the back-end in Rust but since the goal is to primarily only use the standard library, the Rust standard library is very low level which would make "basic" functionality a lot more involved. Given this, I'm wrote the back-end in Go and I'm very happy with the end result.
* Back-end saves the dataset to a `data.json` file on disk in `/front-end/json/`. The location was more to do to with what it was allowed to see in the front-end.
* I originally used [`npx create-next-app@latest`](https://react.dev/learn/start-a-new-react-project#nextjs-pages-router) to create the React app however then I ran into issues with Next being able to access the local file due to server component limitations and as mentioned in our call, the specifics of front-end frameworks is not my strongest area so through troubleshooting this, I found that Remix was a better framework/DX for this type of app and used [`npx create-remix`](https://react.dev/learn/start-a-new-react-project#remix) to get started again.
* The front-end is admittedly half baked. It is able to see the file and shows the individual review records as rows in a table but I was unable to get two requirements to work: (1) the button to trigger the reload of more data. The button wasn't responding to the click handler which I had set to call a function. And (2) the JSON array from the file was deserialized into a `JsonifyObject` but I was unable to filter on this because it resulted in an undefined value.
* React is not my strong suit so I took a pause on these. I do believe I could pick it up again with some mentorship but my experience has mostly been back-end in Rust/Go and infra/DevOps over the past few years.
* My goal for this was that the data would be pulled from the file, then filtered to the last 48 hours (had to adapt to roughly 72 hours) and then the button would be able to call the endpoint and trigger a table reload with the new data.

# Dev Setup

Clone the repo and run `$ cd back-end && go run *.go` from one terminal session and `$ cd front-end && npm run dev` in another. The back-end app will be available on http://localhost:5050, front-end app on http://localhost:3000. This is vaguely clunky but service orchestration was not part of the problem statement.

# Versions Used

* Go 1.21.7
* Node 20.11.0
