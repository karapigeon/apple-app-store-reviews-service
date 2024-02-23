# Notes

* My current machine is very lacking so using a GitHub Codespace to work on this. I configured the dev container to enable the Docker-in-Docker feature.
* I used [`npx create-next-app@latest`](https://react.dev/learn/start-a-new-react-project#nextjs-pages-router) to create the React app. Front-end is not my area of expertise so this was not an opinionated decision but the easy-DX starter.
* I was originally going to write the back-end in Rust but then going through the document, it mentioned primarily only using standard library APIs and the Rust standard library is very low level which would make a lot of "basic" functionality more involved.
* Given this, I'm going to write the back-end in Go because the Go standard library is very good and it would be the same amount of _floundering_ than if I were to use a different language. I've worked with TypeScript, JavaScript, Go, and Rust in my last couple roles so might as well use the stack preferred by the exercise.